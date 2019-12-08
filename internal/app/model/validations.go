package model

import validation "github.com/go-ozzo/ozzo-validation/v3"

func requiredIf(expression bool) validation.RuleFunc {
	return func(value interface{}) error {
		if expression {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
