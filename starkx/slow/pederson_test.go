package starkx

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
	"github.com/phoebetronic/dydxv3/starkx/point"
)

func Test_Starkx_Pederson(t *testing.T) {
	testCases := []struct {
		a *big.Int
		b *big.Int
		p *point.Point
	}{
		// Case 0
		{
			a: hexutil.MustDecodeBig("0x3d937c035c878245caf64531a5756109c53068da139362728feb561405371cb"),
			b: hexutil.MustDecodeBig("0x208a0a10250e382e1e4bbe2880906c2791bf6275695e02fbbc6aeff9cd8b31a"),
			p: &point.Point{
				X: hexutil.MustDecodeBig("0x30e480bed5fe53fa909cc0f8c4d99b8f9f2c016be4c41e13a4848797979c662"),
				Y: hexutil.MustDecodeBig("0x42f5b6cf886f379c81357293f7a109eeb4728650b2cec0389da1618797568c2"),
			},
		},
		// Case 1
		{
			a: hexutil.MustDecodeBig("0x58f580910a6ca59b28927c08fe6c43e2e303ca384badc365795fc645d479d45"),
			b: hexutil.MustDecodeBig("0x78734f65a067be9bdb39de18434d71e79f7b6466a4b66bbd979ab9e7515fe0b"),
			p: &point.Point{
				X: hexutil.MustDecodeBig("0x68cc0b76cddd1dd4ed2301ada9b7c872b23875d5ff837b3a87993e0d9996b87"),
				Y: hexutil.MustDecodeBig("0x224a3b72f16b13af452f5282ecee5bed77f08f0728a71193949444f7e42eda4"),
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var p *point.Point
			{
				p = Pederson(tc.a, tc.b)
			}

			if p.X.Cmp(tc.p.X) != 0 {
				t.Fatalf("x\n\n%s\n", cmp.Diff(hexutil.Encode(tc.p.X.Bytes()), hexutil.Encode(p.X.Bytes())))
			}
			if p.Y.Cmp(tc.p.Y) != 0 {
				t.Fatalf("y\n\n%s\n", cmp.Diff(hexutil.Encode(tc.p.Y.Bytes()), hexutil.Encode(p.Y.Bytes())))
			}
		})
	}
}
