package point

import (
	"math/big"
)

func Mod(n *big.Int, m *big.Int, p *big.Int) *big.Int {
	x, _, _ := GCD(m, p)
	return big.NewInt(0).Mod(big.NewInt(0).Mul(n, x), p)
}
