package position

import (
	"encoding/json"

	"github.com/phoebetron/dydxv3/client/request"
	"github.com/xh3b4sd/tracer"
)

func (p *P) Get(req Request) (Response, error) {
	var err error

	var byt []byte
	{
		byt, err = p.req.Get("positions", request.Values(req))
		if err != nil {
			return Response{}, tracer.Mask(err)
		}
	}

	var res Response
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return Response{}, tracer.Mask(err)
		}
	}

	return res, nil
}
