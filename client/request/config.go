package request

import (
	"net/http"

	"github.com/phoebetron/dydxv3/client/secret"
)

type Config struct {
	Cli *http.Client
	Pri bool
	Sec secret.Secret
	Tes bool
}

func (c Config) Ensure() Config {
	if c.Cli == nil {
		c.Cli = &http.Client{}
	}

	return c
}

func (c Config) Verify() {
	if c.Cli == nil {
		panic("Config.Cli must not be empty")
	}

	if c.Pri {
		c.Sec.Verify()
	}
}
