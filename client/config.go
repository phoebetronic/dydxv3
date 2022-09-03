package client

import "github.com/phoebetron/dydxv3/client/secret"

type Config struct {
	Sec secret.Secret
	Tes bool
}

func (c Config) Verify() {
	{
		c.Sec.Verify()
	}
}
