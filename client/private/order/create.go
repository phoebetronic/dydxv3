package order

import (
	"encoding/json"
	"strings"

	"github.com/xh3b4sd/tracer"
	"github.com/yanue/starkex"
)

func (o *O) Create(req CreateRequest) (CreateResponse, error) {
	var err error

	{
		req.Signature, err = starkex.OrderSign(strings.TrimPrefix(o.sec.StkPrk, "0x"), o.sigpar(req))
		if err != nil {
			return CreateResponse{}, tracer.Mask(err)
		}
	}

	var byt []byte
	{
		byt, err = o.req.Post("orders", req)
		if err != nil {
			return CreateResponse{}, tracer.Mask(err)
		}
	}

	var res CreateResponse
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return CreateResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}

func (o *O) sigpar(req CreateRequest) starkex.OrderSignParam {
	var net int
	if o.tes {
		net = 3
	} else {
		net = 1
	}

	return starkex.OrderSignParam{
		NetworkId:  net,
		PositionId: req.Position,
		Market:     req.Market,
		Side:       req.Side,
		HumanSize:  req.Size,
		HumanPrice: req.Price,
		LimitFee:   req.LimitFee,
		ClientId:   req.ClientId,
		Expiration: req.Expiration.UTC().Format("2006-01-02T15:04:05.000Z"),
	}
}
