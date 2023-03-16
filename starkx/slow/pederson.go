package starkx

import (
	"encoding/json"
	"math/big"

	"github.com/phoebetronic/dydxv3/starkx/config"
	"github.com/phoebetronic/dydxv3/starkx/point"
)

var (
	nebits = 252
	params config.Config
)

func init() {
	err := json.Unmarshal([]byte(config.JSON), &params)
	if err != nil {
		panic(err)
	}
}

func Pederson(e ...*big.Int) *point.Point {
	p := &point.Point{
		X: params.ConstantPoints[0][0],
		Y: params.ConstantPoints[0][1],
	}

	for i, x := range e {
		l := params.ConstantPoints[2+i*nebits : 2+(i+1)*nebits]
		if len(l) != nebits {
			panic("invalid point list")
		}

		for _, c := range l {
			if p.X.Cmp(c[0]) == 0 {
				panic("unhashable input")
			}

			if big.NewInt(0).And(x, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
				p = point.Add(p, &point.Point{X: c[0], Y: c[1]}, params.FieldPrime)
			}

			{
				x.Rsh(x, 1)
			}
		}

		if x.Cmp(big.NewInt(0)) != 0 {
			panic("element must be zero")
		}
	}

	return p
}
