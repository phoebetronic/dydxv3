package account

import (
	"time"
)

type Account struct {
	PositionId         int64     `json:"positionId,string"`
	ID                 string    `json:"id"`
	StarkKey           string    `json:"starkKey"`
	Equity             string    `json:"equity"`
	FreeCollateral     string    `json:"freeCollateral"`
	QuoteBalance       string    `json:"quoteBalance"`
	PendingDeposits    string    `json:"pendingDeposits"`
	PendingWithdrawals string    `json:"pendingWithdrawals"`
	AccountNumber      string    `json:"accountNumber"`
	CreatedAt          time.Time `json:"createdAt"`
}

type Request struct {
	// Address is the Ethereum address of the underlying dYdX account to lookup.
	Address string `json:"address"`
}

type Response struct {
	Account Account `json:"account"`
}
