package clienv

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidEnvironmentError = &tracer.Error{
	Kind: "invalidEnvironmentError",
}

func IsInvalidEnvironment(err error) bool {
	return errors.Is(err, invalidEnvironmentError)
}
