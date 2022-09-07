package market

import (
	"github.com/phoebetron/dydxv3/client/request"
)

type Config struct {
	Req *request.Request
}

func (c Config) Verify() {
	if c.Req == nil {
		panic("Config.Req must not be empty")
	}
}
