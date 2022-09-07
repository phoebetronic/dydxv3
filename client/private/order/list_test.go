package order

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/dydxv3/client/request"
)

func Test_Client_Private_Order_List_Request(t *testing.T) {
	req := ListRequest{
		Market:             "ETH-USD",
		ReturnLatestOrders: true,
	}

	cur := request.Values(req).Encode()
	des := "Market=ETH-USD&ReturnLatestOrders=true"

	if cur != des {
		t.Fatalf("\n\n%s\n", cmp.Diff(des, cur))
	}
}
