package user

import "time"

type ApiKey struct {
	Key        string `json:"key"`
	Passphrase string `json:"passphrase"`
	Secret     string `json:"secret"`
}

type User struct {
	EthereumAddress string `json:"ethereumAddress"`
	IsRegistered    bool   `json:"isRegistered"`
	MakerFeeRate    string `json:"makerFeeRate"`
	TakerFeeRate    string `json:"takerFeeRate"`
}

type Account struct {
	StarkKey  string    `json:"starkKey"`
	CreatedAt time.Time `json:"createdAt"`
}

type Request struct {
	StarkKey            string `json:"starkKey"`
	StarkKeyYCoordinate string `json:"starkKeyYCoordinate"`
}
