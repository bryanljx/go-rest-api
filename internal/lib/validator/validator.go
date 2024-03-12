package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func New() Validator {
	validate := validator.New()

	// Using the names which have been specified for JSON representations of structs, rather than normal Go field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return Validator{
		validator: validate,
	}
}

func (v *Validator) ValidateStruct(params any) error {
	return v.validator.Struct(params)
}

func (v *Validator) ValidateField(params any, checks string) error {
	return v.validator.Var(params, checks)
}
