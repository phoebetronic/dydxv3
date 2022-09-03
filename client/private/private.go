package private

import (
	"github.com/phoebetron/dydxv3/client/private/account"
	"github.com/phoebetron/dydxv3/client/private/fill"
	"github.com/phoebetron/dydxv3/client/private/order"
	"github.com/phoebetron/dydxv3/client/private/position"
	"github.com/phoebetron/dydxv3/client/private/user"
)

type Private struct {
	Acc *account.A
	Fil *fill.F
	Ord *order.O
	Pos *position.P
	Use *user.U
}

func New(con Config) *Private {
	{
		con.Verify()
	}

	var acc *account.A
	{
		acc = account.New(account.Config{
			Req: con.Req,
		})
	}

	var fil *fill.F
	{
		fil = fill.New(fill.Config{
			Req: con.Req,
		})
	}

	var ord *order.O
	{
		ord = order.New(order.Config{
			Req: con.Req,
			Sec: con.Sec,
			Tes: con.Tes,
		})
	}

	var pos *position.P
	{
		pos = position.New(position.Config{
			Req: con.Req,
		})
	}

	var use *user.U
	{
		use = user.New(user.Config{
			Req: con.Req,
		})
	}

	return &Private{
		Acc: acc,
		Fil: fil,
		Ord: ord,
		Pos: pos,
		Use: use,
	}
}
