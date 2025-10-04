package auth

import (
	"github.com/amirtavakolian/quiz-game/param/authparams"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type VerifyValidator struct {
}

func NewVerifyValidator() VerifyValidator {
	return VerifyValidator{}
}

func (a VerifyValidator) Verify(params *authparams.VerifyParam) error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^09(0[1-5]|1\\d|2[0-2]|3\\d|9\\d)\\d{6}$"))),
		validation.Field(&params.OTPCode, validation.Required, validation.Length(6, 6)),
	)
}
