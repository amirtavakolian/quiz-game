package profilehandler

import (
	"github.com/amirtavakolian/quiz-game/param/profileparams"
	"github.com/amirtavakolian/quiz-game/service/profileservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileHandler struct {
	ProfileSvc profileservice.Profile
}

func NewProfileHandler(ProfileService profileservice.Profile) ProfileHandler {
	return ProfileHandler{
		ProfileSvc: ProfileService,
	}
}

func (p ProfileHandler) Update(c echo.Context) error {
	var storeProfileParams profileparams.UpdateProfile

	if err := c.Bind(&storeProfileParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := p.ProfileSvc.Update(c, storeProfileParams)

	return c.JSON(result.StatusCode, result)
}
