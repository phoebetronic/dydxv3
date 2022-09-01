package signer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/phoebetron/dydxv3/starkx/private"
)

// Keyp derives the Stark key pair for the underlying wallet.
//
//     https://docs.dydx.exchange/#derive-starkkey
//
func (s *Signer) Keyp() *private.Key {
	sig, _ := new(big.Int).SetString(s.Sign("dYdX STARK Key"), 0)

	has := solsha3.SoliditySHA3(
		math.PaddedBigBytes(sig, 32),
	)

	return private.New(new(big.Int).Rsh(new(big.Int).SetBytes(has), 5))
}
