package client

import (
	"net/http"

	"github.com/phoebetron/dydxv3/client/private"
	"github.com/phoebetron/dydxv3/client/request"
	"github.com/phoebetron/dydxv3/client/secret"
)

type Client struct {
	Pri *private.Private
	sec secret.Secret
}

func New(con Config) *Client {
	{
		con.Verify()
	}

	var req *request.Request
	{
		req = request.New(request.Config{
			Cli: &http.Client{},
			Sec: con.Sec,
			Tes: con.Tes,
		})
	}

	var pri *private.Private
	{
		pri = private.New(private.Config{
			Req: req,
			Sec: con.Sec,
			Tes: con.Tes,
		})
	}

	return &Client{
		Pri: pri,
		sec: con.Sec,
	}
}
