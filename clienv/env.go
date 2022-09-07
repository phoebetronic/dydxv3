package clienv

const (
	Dydxv3ApiKey                    = "DYDXV3_API_KEY"
	Dydxv3ApiPassphrase             = "DYDXV3_API_PASSPHRASE"
	Dydxv3ApiSecret                 = "DYDXV3_API_SECRET"
	Dydxv3EthAddress                = "DYDXV3_ETH_ADDRESS"
	Dydxv3StarkPrivateKey           = "DYDXV3_STARK_PRIVATE_KEY"
	Dydxv3StarkPublicKey            = "DYDXV3_STARK_PUBLIC_KEY"
	Dydxv3StarkPublicKeyYCoordinate = "DYDXV3_STARK_PUBLIC_KEY_Y_COORDINATE"
)

type Env struct {
	ApiKey                    string
	ApiPassphrase             string
	ApiSecret                 string
	EthAddress                string
	StarkPrivateKey           string
	StarkPublicKey            string
	StarkPublicKeyYCoordinate string
}
