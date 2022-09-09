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

func Values(str interface{}) url.Values {
	var val reflect.Value
	{
		val = reflect.ValueOf(str)
	}

	var typ reflect.Type
	{
		typ = val.Type()
	}

	dic := url.Values{}
	for i := 0; i < val.NumField(); i++ {
		if !val.Field(i).IsZero() && typ.Field(i).Tag.Get("json") != "-" {
			dic.Set(typ.Field(i).Tag.Get("json"), fmt.Sprintf("%v", val.Field(i).Interface()))
		}
	}

	return dic
}
