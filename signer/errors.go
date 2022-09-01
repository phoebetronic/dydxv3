package signer

// Errors is used to identify error responses during the onboarding process. A
// typical error may be an unfunded wallet that has just been created.
//
//     {"errors":[{"msg":"User wallet has no transactions, Ethereum or USDC"}]}
//
type Errors struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Msg string `json:"msg"`
}
