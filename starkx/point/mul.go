package point

import (
	"math/big"
)

func Mul(p *big.Int, g *Point, a *big.Int, f *big.Int) *Point {
	if p.Cmp(big.NewInt(1)) == 0 {
		return g
	}

	if big.NewInt(0).Mod(p, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return Mul(
			big.NewInt(0).Div(p, big.NewInt(2)),
			Dbl(g, a, f),
			a,
			f,
		)
	}

	return Add(
		Mul(
			big.NewInt(0).Sub(p, big.NewInt(1)),
			g,
			a,
			f,
		),
		g,
		f,
	)
}
