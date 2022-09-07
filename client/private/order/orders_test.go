package order

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Client_Private_Order_Orders_Pri(t *testing.T) {
	testCases := []struct {
		ord Orders
		pri float32
	}{
		// case 0
		{
			ord: Orders{},
			pri: 0,
		},
		// case 1
		{
			ord: Orders{
				{
					Price: "2996",
					Size:  "0.3",
				},
				{
					Price: "3005",
					Size:  "0.1",
				},
			},
			pri: 2998.25,
		},
		// case 2
		{
			ord: Orders{
				{
					Price: "3000",
					Size:  "0.5",
				},
				{
					Price: "2998",
					Size:  "0.1",
				},
				{
					Price: "3008",
					Size:  "0.2",
				},
			},
			pri: 3001.75,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var pri float32
			{
				pri = tc.ord.Pri()
			}

			if tc.pri != pri {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.pri, pri))
			}
		})
	}
}
