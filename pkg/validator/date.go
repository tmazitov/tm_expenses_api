package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func dateValidation(fl validator.FieldLevel) bool {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	_, err := time.Parse(fl.Param(), s)
	return err == nil
}
