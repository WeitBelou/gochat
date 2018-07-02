package api

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type defaultValidator struct {
	*validator.Validate
}

func (v *defaultValidator) Engine() interface{} {
	return v.Validate
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := v.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func NewValidator() *defaultValidator {
	v := &defaultValidator{
		Validate: validator.New(),
	}

	v.SetTagName("binding")
	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}
