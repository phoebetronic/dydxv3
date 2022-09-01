package point

import "math/big"

func GCD(a *big.Int, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	x := big.NewInt(0)
	y := big.NewInt(0)
	g := big.NewInt(0)

	if a.Cmp(x) == 0 && b.Cmp(y) == 0 {
		y = big.NewInt(1)
	} else {
		g = g.GCD(x, y, a, b)
	}

	return x, y, g
}
