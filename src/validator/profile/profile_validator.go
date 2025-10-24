package profilevalidator

import (
	"github.com/amirtavakolian/quiz-game/param/profileparams"
	"github.com/go-ozzo/ozzo-validation/v4"
)

type Profile struct {
}

func NewProfileValidator() Profile {
	return Profile{}
}

func (a Profile) Validate(params profileparams.UpdateProfile) error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.Fullname,
			validation.When(params.Fullname != "",
				validation.Length(3, 250),
			),
		),
		validation.Field(&params.Bio,
			validation.When(params.Bio != "",
				validation.Length(2, 500),
			),
		),
	)
}
