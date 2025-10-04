package auth

import (
	"github.com/amirtavakolian/quiz-game/param/authparams"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type AuthValidator struct {
}

func NewAuthValidator() AuthValidator {
	return AuthValidator{}
}

func (a AuthValidator) Authenticate(params authparams.RegisterParam) error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^09\d{9}$`))),
	)
}

func (a AuthValidator) Verify(params *authparams.VerifyParam) error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^09(0[1-5]|1\\d|2[0-2]|3\\d|9\\d)\\d{6}$"))),
		validation.Field(&params.OTPCode, validation.Required, validation.Length(6, 6)),
	)
}
