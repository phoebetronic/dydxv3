package secret

type Secret struct {
	StkPrk string
	StkPux string
	StkPuy string
	EthAdd string
	ApiKey string
	ApiPas string
	ApiSec string
}

func (s Secret) Verify() {
	if s.StkPrk == "" {
		panic("Secret.StkPrk must not be empty")
	}
	if s.StkPux == "" {
		panic("Secret.StkPux must not be empty")
	}
	if s.StkPuy == "" {
		panic("Secret.StkPuy must not be empty")
	}
	if s.EthAdd == "" {
		panic("Secret.EthAdd must not be empty")
	}
	if s.ApiKey == "" {
		panic("Secret.ApiKey must not be empty")
	}
	if s.ApiPas == "" {
		panic("Secret.ApiPas must not be empty")
	}
	if s.ApiSec == "" {
		panic("Secret.ApiSec must not be empty")
	}
}
