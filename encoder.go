package trongrid

import (
	"reflect"
	"time"

	"github.com/gorilla/schema"
)

func NewEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()
	encoder.RegisterEncoder(time.Time{}, func(v reflect.Value) string {
		t, ok := v.Interface().(time.Time)
		if !ok {
			return ""
		}

		return t.Format(layout)
	})
	encoder.SetAliasTag("url")

	return encoder
}
