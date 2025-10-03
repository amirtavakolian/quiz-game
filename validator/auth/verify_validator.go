package auth

import (

"github.com/amirtavakolian/quiz-game/param/authparams"
)

type VerifyValidator struct {
}

func NewVerifyValidator() VerifyValidator {
	return VerifyValidator{}
}

func (u VerifyValidator) Validate(param authparams.VerifyParam) (bool, map[string][]string) {
	errorsList := make(map[string][]string)
	status := true

	if len(param.PhoneNumber) != 11 || param.PhoneNumber[:2] != "09" {
		errorsList["errors"] = append(errorsList["errors"], "phone number must be 11 numbers and start with 09")
		status = false
	}

	if len(param.OTPCode) != 6 {
		errorsList["errors"] = append(errorsList["errors"], "OTP code must be 6 digit")
		status = false
	}

	return status, errorsList
}
