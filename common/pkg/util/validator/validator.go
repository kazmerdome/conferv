package validator

import (
	v "github.com/go-playground/validator"
)

var validatorCache *v.Validate

type structValidator struct {
	Validate *v.Validate
}

func NewStructValidator() *structValidator {
	if validatorCache == nil {
		return &structValidator{
			Validate: v.New(),
		}
	}
	return &structValidator{
		Validate: validatorCache,
	}
}

func ValidateStruct(s interface{}) error {
	vr := NewStructValidator()
	return vr.Validate.Struct(s)
}
