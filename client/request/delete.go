package request

import (
	"net/http"
	"net/url"
)

func (p *Request) Delete(pat string, par url.Values) ([]byte, error) {
	return p.request(http.MethodDelete, Query(pat, par), nil)
}
