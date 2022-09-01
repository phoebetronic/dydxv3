package signer

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common/math"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type Apic struct {
	Key        string
	Passphrase string
	Secret     string
}

// Apic recovers the default API credentials of the underlying wallet.
//
//     https://docs.dydx.exchange/#recover-default-api-credentials
//
func (s *Signer) Apic() Apic {
	sig := s.Sign("dYdX Onboarding")

	rInt, _ := math.MaxBig256.SetString(sig[2:66], 16)
	hashedRBytes := solsha3.SoliditySHA3([]string{"uint256"}, rInt.String())
	secretBytes := hashedRBytes[:30]

	sInt, _ := math.MaxBig256.SetString(sig[66:130], 16)
	hashedSBytes := solsha3.SoliditySHA3([]string{"uint256"}, sInt.String())
	passphraseBytes := hashedSBytes[16:31]

	keyHex := hex.EncodeToString(hashedSBytes[:16])
	keyUuid := strings.Join([]string{
		keyHex[:8],
		keyHex[8:12],
		keyHex[12:16],
		keyHex[16:20],
		keyHex[20:],
	}, "-")

	return Apic{
		Key:        keyUuid,
		Passphrase: base64.URLEncoding.EncodeToString(passphraseBytes),
		Secret:     base64.URLEncoding.EncodeToString(secretBytes),
	}
}
