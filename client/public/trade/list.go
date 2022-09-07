package trade

import (
	"encoding/json"

	"github.com/phoebetron/dydxv3/client/request"
	"github.com/xh3b4sd/tracer"
)

func (t *T) List(req ListRequest) (ListResponse, error) {
	var err error

	var byt []byte
	{
		byt, err = t.req.Get("trades/"+req.Market, request.Values(req))
		if err != nil {
			return ListResponse{}, tracer.Mask(err)
		}
	}

	var res ListResponse
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return ListResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
