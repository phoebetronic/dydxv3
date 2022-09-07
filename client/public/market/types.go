package market

import (
	"strconv"
)

type Market struct {
	Market           string `json:"market"`
	Status           string `json:"status"`
	BaseAsset        string `json:"baseAsset"`
	QuoteAsset       string `json:"quoteAsset"`
	StepSize         string `json:"stepSize"`
	TickSize         string `json:"tickSize"`
	IndexPrice       string `json:"indexPrice"`
	OraclePrice      string `json:"oraclePrice"`
	MinOrderSize     string `json:"minOrderSize"`
	SyntheticAssetId string `json:"syntheticAssetId"`
}

// MBS is the Minimum Bid Size an exchange allows to be traded, e.g. 0.01 ETH.
// MBS is referred to as minOrderSize in the dYdX API docs.
func (m Market) MBS() float32 {
	f, e := strconv.ParseFloat(m.MinOrderSize, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

// MSP is the Minimum Step Price an exchange allows to be traded. Considering
// 2100.40 USD as an initial bid price and 0.1 as MSP would result in the next
// valid bid price of 2100.50. MSP is referred to as tickSize in the dYdX API
// docs.
func (m Market) MSP() float32 {
	f, e := strconv.ParseFloat(m.TickSize, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

// MSS is the Minimum Step Size an exchange allows to be traded. Considering
// 0.021 ETH as an initial bid size and 0.001 as MSS would result in the next
// valid bid size of 0.022. MSS is referred to as stepSize in the dYdX API docs.
func (m Market) MSS() float32 {
	f, e := strconv.ParseFloat(m.StepSize, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

type ListRequest struct {
	Market string `json:"market,omitempty"`
}

type ListResponse struct {
	Markets []Market `json:"markets"`
}
