package fill

import (
	"strconv"
	"time"
)

type Fill struct {
	ID        string    `json:"id"`
	Side      string    `json:"side"`
	Liquidity string    `json:"liquidity"`
	Type      string    `json:"type"`
	Market    string    `json:"market"`
	OrderID   string    `json:"orderId"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Fee       string    `json:"fee"`
	CreatedAt time.Time `json:"createdAt"`
}

func (f Fill) Pri() float32 {
	p, e := strconv.ParseFloat(f.Price, 32)
	if e != nil {
		panic(e)
	}

	return float32(p)
}

func (f Fill) Siz() float32 {
	p, e := strconv.ParseFloat(f.Size, 32)
	if e != nil {
		panic(e)
	}

	return float32(p)
}

type Request struct {
	Market            string `json:"market,omitempty"`
	OrderId           string `json:"order_id,omitempty"`
	Limit             string `json:"limit,omitempty"`
	CreatedBeforeOrAt string `json:"createdBeforeOrAt,omitempty"`
}

type Response struct {
	Fills []Fill `json:"fills"`
}
