package point

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Starkx_Point_GCD(t *testing.T) {
	testCases := []struct {
		a *big.Int
		b *big.Int
		x *big.Int
		y *big.Int
		g *big.Int
	}{
		// Case 0
		{
			a: big.NewInt(0),
			b: big.NewInt(0),
			x: big.NewInt(0),
			y: big.NewInt(1),
			g: big.NewInt(0),
		},
		// Case 1
		{
			a: big.NewInt(2),
			b: big.NewInt(3),
			x: big.NewInt(-1),
			y: big.NewInt(1),
			g: big.NewInt(1),
		},
		// Case 2
		{
			a: big.NewInt(10),
			b: big.NewInt(12),
			x: big.NewInt(-1),
			y: big.NewInt(1),
			g: big.NewInt(2),
		},
		// Case 3
		{
			a: big.NewInt(100),
			b: big.NewInt(2004),
			x: big.NewInt(-20),
			y: big.NewInt(1),
			g: big.NewInt(4),
		},
		// Case 4
		{
			a: big.NewInt(30),
			b: big.NewInt(50),
			x: big.NewInt(2),
			y: big.NewInt(-1),
			g: big.NewInt(10),
		},
		// Case 5
		{
			a: big.NewInt(10),
			b: big.NewInt(0),
			x: big.NewInt(1),
			y: big.NewInt(0),
			g: big.NewInt(10),
		},
		// Case 6
		{
			a: big.NewInt(0),
			b: big.NewInt(10),
			x: big.NewInt(0),
			y: big.NewInt(1),
			g: big.NewInt(10),
		},
		// Case 7
		{
			a: big.NewInt(98),
			b: big.NewInt(56),
			x: big.NewInt(-1),
			y: big.NewInt(2),
			g: big.NewInt(14),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var x *big.Int
			var y *big.Int
			var g *big.Int
			{
				x, y, g = GCD(tc.a, tc.b)
			}

			if x.Cmp(tc.x) != 0 {
				t.Fatalf("x\n\n%s\n", cmp.Diff(tc.x.String(), x.String()))
			}
			if y.Cmp(tc.y) != 0 {
				t.Fatalf("y\n\n%s\n", cmp.Diff(tc.y.String(), y.String()))
			}
			if g.Cmp(tc.g) != 0 {
				t.Fatalf("g\n\n%s\n", cmp.Diff(tc.g.String(), g.String()))
			}
		})
	}
}
