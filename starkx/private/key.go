package private

import (
	"encoding/json"
	"math/big"

	"github.com/phoebetronic/dydxv3/starkx/config"
	"github.com/phoebetronic/dydxv3/starkx/point"
)

var (
	params config.Config
)

func init() {
	err := json.Unmarshal([]byte(config.JSON), &params)
	if err != nil {
		panic(err)
	}
}

// Key provides the Stark private key and the Stark public key X and Y
// coordinates. A string helper like go-ethereum's hexutil may be used to render
// they key's *big.Int types.
//
//	github.com/ethereum/go-ethereum/common/hexutil
//
// Having created key via private.New enables the caller to further use the
// generated key pair in hex representation.
//
//	stark private key                    hexutil.EncodeBig(key.Pri)
//	stark public key                     hexutil.EncodeBig(key.Pub.X)
//	stark public key y coordinate        hexutil.EncodeBig(key.Pub.Y)
type Key struct {
	Pri *big.Int
	Pub *point.Point
}

func New(pri *big.Int) *Key {
	var g *point.Point
	{
		g = &point.Point{
			X: params.ConstantPoints[1][0],
			Y: params.ConstantPoints[1][1],
		}
	}

	return &Key{
		Pri: pri,
		Pub: point.Mul(pri, g, params.Alpha, params.FieldPrime),
	}
}
