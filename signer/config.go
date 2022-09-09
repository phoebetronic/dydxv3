package signer

import (
	"github.com/phoebetron/wallet/pkg/wallet"
)

type Config struct {
	// Tes can be set to true in order to work against dYdX's staging
	// environment on the Ropsten testnet.
	Tes bool
	// Wal provides an ECDSA private key of an Ethereum wallet used for Stark
	// key cryptography and EIP712 typed data signing.
	Wal *wallet.Wallet
}

func (c Config) Verify() {
	if c.Wal == nil {
		panic("Config.Wal must not be empty")
	}
}
