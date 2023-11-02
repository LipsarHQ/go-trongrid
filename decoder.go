package trongrid

import (
	"reflect"
	"time"

	"github.com/gorilla/schema"
)

func NewDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.RegisterConverter(time.Time{}, func(s string) reflect.Value {
		t, err := time.Parse(layout, s)
		if err != nil {
			return reflect.ValueOf("")
		}

		return reflect.ValueOf(t)
	})
	decoder.SetAliasTag("url")
	decoder.ZeroEmpty(true)

	return decoder
}

// func ParseURI(s string, v any) (err error) {
//	var vs url.Values
//
//	if vs, err = url.ParseQuery(s); err != nil {
//		return err
//	}
//
//	if err = decoder.Decode(&v, vs); err != nil {
//		return err
//	}
//
//	return nil
// }.
