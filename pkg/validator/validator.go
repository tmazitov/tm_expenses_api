package validator

import (
	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *validator.Validate
}

func New() *StructValidator {
	validate := validator.New()

	// min value check
	validate.RegisterValidation("decimal_min", minDecimalValidation)
	validate.RegisterValidation("date", dateValidation)

	return &StructValidator{
		validate: validate,
	}
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}
