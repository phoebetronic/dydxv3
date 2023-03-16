package order

import (
	"github.com/phoebetronic/dydxv3/client/request"
	"github.com/phoebetronic/dydxv3/client/secret"
)

type O struct {
	req *request.Request
	sec secret.Secret
	tes bool
}

func New(con Config) *O {
	{
		con.Verify()
	}

	return &O{
		req: con.Req,
		sec: con.Sec,
		tes: con.Tes,
	}
}
