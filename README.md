# dydxv3

Golang implementation of the [StarkEx Pederson Hash Function] used for [dYdX
Exchange Authentication].

[StarkEx Pederson Hash Function]: https://docs.starkware.co/starkex-v4/crypto/pedersen-hash-function
[dYdX Exchange Authentication]: https://docs.dydx.exchange/#authentication



#### register a new user account with the dydx api

```go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/phoebetron/dydxv3/signer"
	"github.com/phoebetron/dydxv3/starkx/private"
	"github.com/phoebetron/wallet/pkg/wallet"
)

func main() {
	// Create or restore your ethereum wallet using the 24 mnemonic word list
	// used to derive the wallet's private key. If you provide no word list a
	// new wallet is created.
	sig := signer.New(signer.Config{
		Wal: wallet.New(wallet.Config{Mne: "tea choral speed ... teddy moon brother"}),
	})

	// Register a new user account for the underlying wallet. Once your user
	// account got created you can further use the client implementation in
	// order to access private endpoints.
	var use signer.User
	{
		use = sig.User(sig.Keyp())
	}

	var key *private.Key
	{
		key = sig.Keyp()
	}

	fmt.Printf("stark private key                %s\n", hexutil.EncodeBig(key.Pri))
	fmt.Printf("stark public key                 %s\n", hexutil.EncodeBig(key.Pub.X))
	fmt.Printf("stark public key y coordinate    %s\n", hexutil.EncodeBig(key.Pub.Y))
	fmt.Printf("ethereum address                 %s\n", use.User.EthereumAddress)
	fmt.Printf("dydx api key                     %s\n", use.ApiKey.Key)
	fmt.Printf("dydx api passphrase              %s\n", use.ApiKey.Passphrase)
	fmt.Printf("dydx api secret                  %s\n", use.ApiKey.Secret)
}
```



#### generate client secrets for dydx api access

```go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/phoebetron/dydxv3/client/secret"
	"github.com/phoebetron/dydxv3/signer"
	"github.com/phoebetron/dydxv3/starkx/private"
	"github.com/phoebetron/wallet/pkg/wallet"
)

func main() {
	// Create or restore your ethereum wallet using the 24 mnemonic word list
	// used to derive the wallet's private key. If you provide no word list a
	// new wallet is created.
	sig := signer.New(signer.Config{
		Wal: wallet.New(wallet.Config{Mne: "tea choral speed ... teddy moon brother"}),
	})

	// Recover the dydx default api credentials for the underlying wallet.
	var cre signer.Apic
	{
		cre = sig.Apic()
	}

	// Derive the stark key pair for the underlying wallet.
	var key *private.Key
	{
		key = sig.Keyp()
	}

	// Create the secret configuration required to use the dydx client
	// implementation.
	var sec secret.Secret
	{
		sec = secret.Secret{
			StkPrk: hexutil.EncodeBig(key.Pri),
			StkPux: hexutil.EncodeBig(key.Pub.X),
			StkPuy: hexutil.EncodeBig(key.Pub.Y),
			EthAdd: sig.Wal.Add.String(),
			ApiKey: cre.Key,
			ApiPas: cre.Passphrase,
			ApiSec: cre.Secret,
		}
	}

	fmt.Printf("%#v\n", sec)
}
```



#### access private endpoints of the dydx api

```go
package main

import (
	"fmt"

	"github.com/phoebetron/dydxv3/client"
	"github.com/phoebetron/dydxv3/client/private/account"
	"github.com/phoebetron/dydxv3/client/secret"
)

func main() {
	var cli *client.Client
	{
		cli = client.New(client.Config{
			Sec: secret.Secret{
				StkPrk: "stark private key",
				StkPux: "stark public key",
				StkPuy: "stark public key y coordinate",
				EthAdd: "ethereum address",
				ApiKey: "dydx api key",
				ApiPas: "dydx api passphrase",
				ApiSec: "dydx api secret",
			},
		})
	}

	acc, err := cli.Pri.Acc.Get(account.Request{Address: "ethereum address"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", acc)
}
```
