package user

import (
	"github.com/phoebetronic/dydxv3/client/request"
)

type U struct {
	req *request.Request
}

func New(con Config) *U {
	{
		con.Verify()
	}

	return &U{
		req: con.Req,
	}
}
