package signer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/phoebetron/dydxv3/starkx/private"
	"github.com/phoebetron/dydxv3/user"
	"github.com/xh3b4sd/tracer"
)

type User struct {
	ApiKey  user.ApiKey  `json:"apiKey"`
	User    user.User    `json:"user"`
	Account user.Account `json:"account"`
}

// User let's the caller onboard to the dYdX exchange by creating a new account
// for the underlying wallet.
//
//     signer.User(signer.Keyp())
//
func (s *Signer) User(key *private.Key) User {
	var err error

	var use User
	{
		use, err = s.post(key)
		if err != nil {
			panic(err)
		}
	}

	return use
}

func (s *Signer) endp() string {
	var end string
	if s.Tes {
		end = "https://api.stage.dydx.exchange"
	} else {
		end = "https://api.dydx.exchange"
	}

	return fmt.Sprintf("%s/%s/%s", end, "v3", "onboarding")
}

func (s *Signer) post(key *private.Key) (User, error) {
	var err error

	var req *http.Request
	{
		req, err = http.NewRequest("POST", s.endp(), s.body(key))
		if err != nil {
			return User{}, tracer.Mask(err)
		}
	}

	x := s.Sign("dYdX Onboarding")

	{
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("DYDX-SIGNATURE", x)
		req.Header.Set("DYDX-ETHEREUM-ADDRESS", s.Wal.Add.String())
	}

	var res *http.Response
	{
		res, err = s.Cli.Do(req)
		if err != nil {
			return User{}, tracer.Mask(err)
		}
		defer res.Body.Close()
	}

	var bod []byte
	{
		bod, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return User{}, tracer.Mask(err)
		}
	}

	var lis Errors
	{
		err = json.Unmarshal(bod, &lis)
		if err != nil {
			return User{}, tracer.Mask(err)
		}
	}

	if len(lis.Errors) != 0 {
		return User{}, tracer.Mask(fmt.Errorf(lis.Errors[0].Msg))
	}

	var use User
	{
		err = json.Unmarshal(bod, &use)
		if err != nil {
			return User{}, tracer.Mask(err)
		}
	}

	return use, nil
}

func (s *Signer) body(key *private.Key) io.Reader {
	var err error

	var req user.Request
	{
		req = user.Request{
			StarkKey:            hexutil.EncodeBig(key.Pub.X),
			StarkKeyYCoordinate: hexutil.EncodeBig(key.Pub.Y),
		}
	}

	var byt []byte
	{
		byt, err = json.Marshal(req)
		if err != nil {
			panic(err)
		}
	}

	return bytes.NewBuffer(byt)
}
