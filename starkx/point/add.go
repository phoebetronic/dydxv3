package point

import "math/big"

func Add(a *Point, b *Point, p *big.Int) *Point {
	m := Mod(
		big.NewInt(0).Sub(a.Y, b.Y),
		big.NewInt(0).Sub(a.X, b.X),
		p,
	)

	x := big.NewInt(0).Mod(
		big.NewInt(0).Sub(
			big.NewInt(0).Sub(
				big.NewInt(0).Mul(m, m),
				a.X,
			),
			b.X,
		),
		p,
	)

	y := big.NewInt(0).Mod(
		big.NewInt(0).Sub(
			big.NewInt(0).Mul(
				m,
				big.NewInt(0).Sub(a.X, x),
			),
			a.Y,
		),
		p,
	)

	return &Point{X: x, Y: y}
}
