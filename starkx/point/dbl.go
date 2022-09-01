package point

import (
	"math/big"
)

func Dbl(g *Point, a *big.Int, f *big.Int) *Point {
	m := Mod(
		big.NewInt(0).Add(
			big.NewInt(0).Mul(big.NewInt(3), big.NewInt(0).Mul(g.X, g.X)),
			a,
		),
		big.NewInt(0).Mul(big.NewInt(2), g.Y),
		f,
	)

	x := big.NewInt(0).Mod(
		big.NewInt(0).Sub(
			big.NewInt(0).Mul(m, m),
			big.NewInt(0).Mul(big.NewInt(2), g.X),
		),
		f,
	)

	y := big.NewInt(0).Mod(
		big.NewInt(0).Sub(
			big.NewInt(0).Mul(
				m,
				big.NewInt(0).Sub(g.X, x),
			),
			g.Y,
		),
		f,
	)

	return &Point{X: x, Y: y}
}
