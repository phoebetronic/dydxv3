package signer

import (
	"net/http"

	"github.com/ethersphere/bee/pkg/crypto"
	"github.com/phoebetron/wallet"
)

// Signer implements all required primitives for dYdX authentication.
//
//     https://docs.dydx.exchange/#authentication
//
type Signer struct {
	Cli *http.Client
	Sig crypto.Signer
	Tes bool
	Wal *wallet.Wallet
}

func New(con Config) *Signer {
	{
		con.Verify()
	}

	return &Signer{
		Cli: &http.Client{},
		Sig: crypto.NewDefaultSigner(con.Wal.Pri),
		Tes: con.Tes,
		Wal: con.Wal,
	}
}
