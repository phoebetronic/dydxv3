package account

import (
	"encoding/json"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/xh3b4sd/tracer"
)

func (a *A) Get(req Request) (Response, error) {
	var err error

	var acc uuid.UUID
	{
		acc = uuid.NewV5(a.nam, uuid.NewV5(a.nam, strings.ToLower(req.Address)).String()+"0")
	}

	var byt []byte
	{
		byt, err = a.req.Get("accounts/"+acc.String(), nil)
		if err != nil {
			return Response{}, tracer.Mask(err)
		}
	}

	var res Response
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return Response{}, tracer.Mask(err)
		}
	}

	return res, nil
}
