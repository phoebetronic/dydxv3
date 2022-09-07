package public

import (
	"github.com/phoebetron/dydxv3/client/public/market"
	"github.com/phoebetron/dydxv3/client/public/trade"
)

type Public struct {
	Mar *market.M
	Tra *trade.T
}

func New(con Config) *Public {
	{
		con.Verify()
	}

	var mar *market.M
	{
		mar = market.New(market.Config{
			Req: con.Req,
		})
	}

	var tra *trade.T
	{
		tra = trade.New(trade.Config{
			Req: con.Req,
		})
	}

	return &Public{
		Mar: mar,
		Tra: tra,
	}
}
