package orderbook

import (
	"encoding/json"

	"github.com/xh3b4sd/tracer"
)

func (o *O) Get(req GetRequest) (GetResponse, error) {
	var err error

	var byt []byte
	{
		byt, err = o.req.Get("orderbook/"+req.Market, nil)
		if err != nil {
			return GetResponse{}, tracer.Mask(err)
		}
	}

	var res GetResponse
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return GetResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
