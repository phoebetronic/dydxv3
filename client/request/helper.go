package request

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

var namespace = Must(FromString("0f9da948-a6fb-4c45-9edc-4685c3f3317d"))

func getUserId(address string) string {
	return uuid.NewV5(namespace, address).String()
}

func GetAccountId(address string) string {
	return uuid.NewV5(namespace, getUserId(strings.ToLower(address))+strconv.Itoa(0)).String()
}

func FromString(input string) (u uuid.UUID, err error) {
	err = u.UnmarshalText([]byte(input))
	return
}

func Must(u uuid.UUID, err error) uuid.UUID {
	if err != nil {
		panic(err)
	}
	return u
}

func GenerateQueryPath(url string, par url.Values) string {
	if len(par) == 0 {
		return url
	}

	return fmt.Sprintf("%s?%s", url, par.Encode())
}

func ExpireAfter(duration time.Duration) string {
	return time.Now().Add(duration).UTC().Format("2006-01-02T15:04:05.999Z")
}

func Values(val interface{}) url.Values {
	var ele reflect.Value
	{
		ele = reflect.ValueOf(val).Elem()
	}

	var typ reflect.Type
	{
		typ = ele.Type()
	}

	var dic url.Values
	for i := 0; i < ele.NumField(); i++ {
		dic.Set(typ.Field(i).Name, fmt.Sprint(ele.Field(i)))
	}

	return dic
}
