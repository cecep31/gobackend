package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Valid *validator.Validate

func SetupValidate() {
	Valid = validator.New()
}

func ValidateThis(data interface{}) map[string]string {
	if errvalidate := Valid.Struct(data); errvalidate != nil {
		errMsgs := make(map[string]string)
		for _, err := range errvalidate.(validator.ValidationErrors) {
			errMsgs[err.Field()] = fmt.Sprintf("Field validation for '%s' failed on the '%s'", err.Field(), err.Tag())
		}

		return errMsgs
	}
	return nil
}
