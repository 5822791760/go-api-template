package helpers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func Decode(r *http.Request, dst interface{}) errs.ErrRenderer {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return errs.NewErr(err, http.StatusInternalServerError)
	}

	if err := validate.Struct(dst); err != nil {
		return errs.NewErr(err, http.StatusBadRequest)
	}

	return nil
}

func init() {
	validate.RegisterCustomTypeFunc(validatePointer, (*string)(nil), (*int)(nil), (*int32)(nil), (*int32)(nil), (*float32)(nil), (*float64)(nil), (*bool)(nil))
}

func validatePointer(field reflect.Value) interface{} {
	if field.IsNil() {
		return nil
	}
	return field.Elem().Interface()
}
