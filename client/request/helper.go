package request

import (
	"fmt"
	"net/url"
	"reflect"
)

func Query(url string, par url.Values) string {
	if len(par) == 0 {
		return url
	}

	return fmt.Sprintf("%s?%s", url, par.Encode())
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
