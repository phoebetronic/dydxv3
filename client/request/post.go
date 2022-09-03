package request

import (
	"net/http"
)

func (r *Request) Post(pat string, dat interface{}) ([]byte, error) {
	return r.request(http.MethodPost, pat, dat)
}
