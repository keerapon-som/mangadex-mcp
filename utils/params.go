package utils

import (
	"net/url"
	"reflect"
	"strconv"
)

func ToQueryParams[T any](input T) url.Values {
	v := url.Values{}
	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)
		tag := fieldType.Tag.Get("json")

		if tag == "" || tag == "-" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if val := field.String(); val != "" {
				v.Set(tag, val)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val := field.Int(); val != 0 {
				v.Set(tag, strconv.FormatInt(val, 10))
			}
		case reflect.Slice:
			// If it's a slice of strings, add each with "key[]"
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.String {
					v.Add(tag+"[]", elem.String())
				}
			}
		}
	}

	return v
}
