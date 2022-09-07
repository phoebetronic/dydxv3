package client

import (
	"net/http"

	"github.com/phoebetron/dydxv3/client/private"
	"github.com/phoebetron/dydxv3/client/public"
	"github.com/phoebetron/dydxv3/client/request"
	"github.com/phoebetron/dydxv3/client/secret"
)

type Client struct {
	Pri *private.Private
	Pub *public.Public
}

func New(con Config) *Client {
	{
		con.Verify()
	}

	var pri *private.Private
	if con.Sec != (secret.Secret{}) {
		pri = private.New(private.Config{
			Req: request.New(request.Config{
				Cli: &http.Client{},
				Pri: true,
				Sec: con.Sec,
				Tes: con.Tes,
			}),
			Sec: con.Sec,
			Tes: con.Tes,
		})
	}

	var pub *public.Public
	{
		pub = public.New(public.Config{
			Req: request.New(request.Config{
				Cli: &http.Client{},
				Tes: con.Tes,
			}),
			Tes: con.Tes,
		})
	}

	return &Client{
		Pri: pri,
		Pub: pub,
	}
}
