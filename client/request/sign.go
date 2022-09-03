package request

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func (r *Request) sign(pat string, met string, now string, dat []byte) string {
	sec, err := base64.URLEncoding.DecodeString(r.sec.ApiSec)
	if err != nil {
		panic(err)
	}

	has := hmac.New(sha256.New, sec)
	has.Write([]byte(fmt.Sprintf("%s%s/v3/%s%s", now, met, pat, dat)))

	return base64.URLEncoding.EncodeToString(has.Sum(nil))
}
