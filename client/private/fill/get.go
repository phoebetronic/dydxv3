package fill

import (
	"encoding/json"

	"github.com/phoebetron/dydxv3/client/request"
	"github.com/xh3b4sd/tracer"
)

func (f *F) Get(req Request) (Response, error) {
	var err error

	var byt []byte
	{
		byt, err = f.req.Get("fills", request.Values(req))
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
