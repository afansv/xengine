package plugins

import (
	"net/http"
	"reflect"
)

type ReqWithValidation interface {
	Validate() error
}

type ValidationPlugin struct {
}

func (ValidationPlugin) PreBind(_ interface{}, _ *http.Request) error {
	return nil
}

func (ValidationPlugin) PostBind(i interface{}, _ *http.Request) error {
	t := reflect.TypeOf(i)

	isValidationRequired := t.Implements(reflect.TypeOf((*ReqWithValidation)(nil)).Elem())
	if isValidationRequired {
		if err := i.(ReqWithValidation).Validate(); err != nil {
			return err
		}
	}

	return nil
}
