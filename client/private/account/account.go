package account

import (
	"github.com/phoebetronic/dydxv3/client/request"
	uuid "github.com/satori/go.uuid"
)

type A struct {
	nam uuid.UUID
	req *request.Request
}

func New(con Config) *A {
	{
		con.Verify()
	}

	var nam uuid.UUID
	{
		err := nam.UnmarshalText([]byte("0f9da948-a6fb-4c45-9edc-4685c3f3317d"))
		if err != nil {
			panic(err)
		}
	}

	return &A{
		nam: nam,
		req: con.Req,
	}
}
