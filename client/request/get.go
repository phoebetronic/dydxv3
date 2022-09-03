package request

import (
	"net/http"
	"net/url"
)

func (r *Request) Get(pat string, par url.Values) ([]byte, error) {
	return r.request(http.MethodGet, Query(pat, par), nil)
}
