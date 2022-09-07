package trade

import (
	"strconv"
	"time"
)

type Trade struct {
	Side        string    `json:"side"`
	Price       string    `json:"price"`
	Size        string    `json:"size"`
	CreatedAt   time.Time `json:"createdAt"`
	Liquidation bool      `json:"liquidation"`
}

func (t Trade) Pri() float32 {
	f, e := strconv.ParseFloat(t.Price, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

func (t Trade) Siz() float32 {
	f, e := strconv.ParseFloat(t.Size, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

type ListRequest struct {
	Market             string    `json:"-"`
	Limit              int       `json:"limit"`
	StartingBeforeOrAt time.Time `json:"startingBeforeOrAt"`
}

type ListResponse struct {
	Trades []Trade `json:"trades"`
}
