package validator

import (
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/param"
)

type UserValidator struct {
}

func NewUserValidator() UserValidator {
	return UserValidator{}
}

func (u UserValidator) RegisterUserValidate(param param.RegisterParam) (bool, map[string][]string) {
	errorsList := make(map[string][]string)
	status := true

	if len(param.Name) < 3 {
		errorsList["errors"] = append(errorsList["errors"], "name must be more then 3 charecters")
		status = false
	}

	if len(param.Family) < 3 {
		errorsList["errors"] = append(errorsList["errors"], "firstname must be more then 3 characters")
		status = false
	}

	if len(param.PhoneNumber) < 11 || param.PhoneNumber[:2] != "09" {
		errorsList["errors"] = append(errorsList["errors"], "phone number must be 11 numbers and start with 09")
		status = false
	}

	return status, errorsList
}
