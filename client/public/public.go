package public

import (
	"github.com/phoebetronic/dydxv3/client/public/market"
	"github.com/phoebetronic/dydxv3/client/public/orderbook"
	"github.com/phoebetronic/dydxv3/client/public/trade"
)

type Public struct {
	Mar *market.M
	Ord *orderbook.O
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

	var ord *orderbook.O
	{
		ord = orderbook.New(orderbook.Config{
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
		Ord: ord,
		Tra: tra,
	}
}
