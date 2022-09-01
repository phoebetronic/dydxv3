package signer

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	"github.com/go-numb/go-dydx/types"
)

func (s *Signer) Sign(a string) string {
	sig, err := s.Sig.SignTypedData(s.fil(s.mes(a)))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s0%s", s.fix(hexutil.Encode(sig)), "0")
}

func (s *Signer) fil(m *eip712.TypedData) *eip712.TypedData {
	if !s.Tes {
		return m
	}

	{
		m.Types["dYdX"] = []eip712.Type{m.Types["dYdX"][0]}                // drop signon
		m.Domain.ChainId = math.NewHexOrDecimal256(3)                      // set ropsten
		m.Message = eip712.TypedDataMessage{"action": m.Message["action"]} // drop signon
	}

	return m
}

func (s *Signer) fix(m string) string {
	tri := strings.TrimPrefix(m, "0x")
	if len(tri) != 130 {
		panic(fmt.Sprintf("invalid signature length %d: %s", len(tri), tri))
	}

	rs := tri[:128]
	v := tri[128:130]

	if v == "00" {
		return "0x" + rs + "1b"
	}
	if v == "01" {
		return "0x" + rs + "1c"
	}
	if v == "1b" || v == "1c" {
		return "0x" + tri
	}

	panic(fmt.Sprintf("invalid v value: %s", v))
}

func (s *Signer) mes(a string) *eip712.TypedData {
	return &eip712.TypedData{
		Types: eip712.Types{
			"EIP712Domain": []eip712.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"dYdX": []eip712.Type{
				{
					Type: "string",
					Name: "action",
				},
				{
					Type: "string",
					Name: "onlySignOn",
				},
			},
		},
		Domain: eip712.TypedDataDomain{
			Name:    types.Domain,
			Version: types.Version,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "dYdX",
		Message: eip712.TypedDataMessage{
			"action":     a,
			"onlySignOn": "https://trade.dydx.exchange",
		},
	}
}
