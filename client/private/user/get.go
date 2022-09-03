package user

import (
	"encoding/json"

	"github.com/xh3b4sd/tracer"
)

func (u *U) Get() (Response, error) {
	var err error

	var byt []byte
	{
		byt, err = u.req.Get("users", nil)
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
