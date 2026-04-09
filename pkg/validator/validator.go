package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

type StructValidator struct {
	validate *validator.Validate
}

func New() *StructValidator {
	validate := validator.New()

	// "required" equivalent — rejects zero value
	validate.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if d, ok := field.Interface().(decimal.Decimal); ok {
			return d.String()
		}
		return nil
	}, decimal.Decimal{})

	// min value check
	validate.RegisterValidation("decimal_min", func(fl validator.FieldLevel) bool {
		var d decimal.Decimal
		switch v := fl.Field().Interface().(type) {
		case decimal.Decimal:
			d = v
		case string:
			var err error
			d, err = decimal.NewFromString(v)
			if err != nil {
				return false
			}
		default:
			return false
		}
		min, err := decimal.NewFromString(fl.Param())
		if err != nil {
			return false
		}
		return d.GreaterThanOrEqual(min)
	})

	return &StructValidator{
		validate: validate,
	}
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}
