package orderbook

import (
	"strconv"
)

type Order struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

func (o Order) Pri() float32 {
	f, e := strconv.ParseFloat(o.Price, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

func (o Order) Siz() float32 {
	f, e := strconv.ParseFloat(o.Size, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

type GetRequest struct {
	Market string `json:"market,omitempty"`
}

type GetResponse struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

func (g GetResponse) Ask() map[string]string {
	m := map[string]string{}

	for _, x := range g.Asks {
		m[x.Price] = x.Size
	}

	return m
}

func (g GetResponse) Bid() map[string]string {
	m := map[string]string{}

	for _, x := range g.Bids {
		m[x.Price] = x.Size
	}

	return m
}
