/* This file is auto-generated, manual edits in this file will be overwritten! */
package utils

import (
	"reflect"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/samber/lo"
)

func InitialiseFakerGenerators() {
	customTimeGenerator()
}

func customTimeGenerator() {
	_ = faker.AddProvider("utcTime", func(v reflect.Value) (interface{}, error) {
		if v.Kind() == reflect.Ptr {
			return lo.ToPtr(time.Now().UTC()), nil
		}
		return time.Now().UTC(), nil
	})
}