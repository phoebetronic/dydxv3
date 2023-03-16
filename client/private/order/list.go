package order

import (
	"encoding/json"

	"github.com/phoebetronic/dydxv3/client/request"
	"github.com/xh3b4sd/tracer"
)

func (o *O) List(req ListRequest) (ListResponse, error) {
	var err error

	var byt []byte
	{
		byt, err = o.req.Get("orders", request.Values(req))
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
