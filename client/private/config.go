package private

import (
	"github.com/phoebetronic/dydxv3/client/request"
	"github.com/phoebetronic/dydxv3/client/secret"
)

type Config struct {
	Req *request.Request
	Sec secret.Secret
	Tes bool
}

func (c Config) Verify() {
	if c.Req == nil {
		panic("Config.Req must not be empty")
	}

	{
		c.Sec.Verify()
	}
}
