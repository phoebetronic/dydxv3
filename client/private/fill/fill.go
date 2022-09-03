package fill

import (
	"github.com/phoebetron/dydxv3/client/request"
)

type F struct {
	req *request.Request
}

func New(con Config) *F {
	{
		con.Verify()
	}

	return &F{
		req: con.Req,
	}
}
