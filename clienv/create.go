package clienv

import "os"

func Create() Env {
	var e Env

	{
		e.ApiKey = os.Getenv(Dydxv3ApiKey)
		e.ApiPassphrase = os.Getenv(Dydxv3ApiPassphrase)
		e.ApiSecret = os.Getenv(Dydxv3ApiSecret)
		e.EthAddress = os.Getenv(Dydxv3EthAddress)
		e.StarkPrivateKey = os.Getenv(Dydxv3StarkPrivateKey)
		e.StarkPublicKey = os.Getenv(Dydxv3StarkPublicKey)
		e.StarkPublicKeyYCoordinate = os.Getenv(Dydxv3StarkPublicKeyYCoordinate)
	}

	{
		verify(e)
	}

	return e
}
