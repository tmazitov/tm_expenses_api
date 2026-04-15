package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// StructToQuery converts a struct to a URL query string.
// Uses `query` tag to determine the key name, falls back to lowercase field name.
// Fields with zero values are skipped.
//
// Example:
//
//	type Filters struct {
//	    Date  string `query:"date"`
//	    Page  int    `query:"page"`
//	}
//	StructToQuery(Filters{Date: "2026-01-15", Page: 2}) // "date=2026-01-15&page=2"
func StructToQuery(s any) string {
	params := url.Values{}

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)

		if value.IsZero() {
			continue
		}

		key := field.Tag.Get("query")
		if key == "" {
			key = strings.ToLower(field.Name)
		}

		params.Set(key, fmt.Sprintf("%v", value.Interface()))
	}

	return params.Encode()
}
