package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func minDecimalValidation(fl validator.FieldLevel) bool {
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
}
