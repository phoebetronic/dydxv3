package order

import (
	"github.com/xh3b4sd/tracer"
)

func (o *O) Cancel(req CancelRequest) error {
	var err error

	{
		_, err = o.req.Delete("orders/"+req.ID, nil)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
