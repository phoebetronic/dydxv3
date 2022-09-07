package clienv

import "fmt"

func verify(e Env) {
	{
		if e.ApiKey == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3ApiKey))
		}
		if e.ApiPassphrase == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3ApiPassphrase))
		}
		if e.ApiSecret == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3ApiSecret))
		}
		if e.EthAddress == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3EthAddress))
		}
		if e.StarkPrivateKey == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3StarkPrivateKey))
		}
		if e.StarkPublicKey == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3StarkPublicKey))
		}
		if e.StarkPublicKeyYCoordinate == "" {
			panic(fmt.Sprintf("${%s} must not be empty", Dydxv3StarkPublicKeyYCoordinate))
		}
	}
}
