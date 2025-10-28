package profileservice

import (
	"github.com/amirtavakolian/quiz-game/param/profileparams"
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	profilevalidator "github.com/amirtavakolian/quiz-game/validator/profile"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type Profile struct {
	Responser        responser.Response
	Logger           logger.Logger
	ProfileValidator profilevalidator.Profile
	ProfileRepo      repositorycontracts.ProfileRepoContract
}

func NewProfileService(
	responser responser.Response,
	logger logger.Logger,
	profileValidator profilevalidator.Profile,
	profileRepo repositorycontracts.ProfileRepoContract,
) Profile {
	return Profile{
		Responser:        responser,
		Logger:           logger,
		ProfileRepo:      profileRepo,
		ProfileValidator: profileValidator,
	}
}

func (p Profile) Update(c echo.Context, storeProfileParams profileparams.UpdateProfile) responser.Response {
	if err := p.ProfileValidator.Validate(storeProfileParams); err != nil {
		p.Responser.SetStatusCode(401).SetData(err.Error())
	}

	tokenData := c.Get("user").(*jwt.Token)
	storeProfileParams.PhoneNumber = tokenData.Claims.(jwt.MapClaims)["phone_number"].(string)
	storeProfileParams.PlayerID = tokenData.Claims.(jwt.MapClaims)["player_id"].(float64)

	err := p.ProfileRepo.Update(storeProfileParams)

	if err != nil {
		p.Logger.Log().Error("Profile-Store", zap.Error(err))
		return p.Responser.SetStatusCode(http.StatusInternalServerError).SetMessage("Internal Error").Build()
	}

	return p.Responser.SetStatusCode(http.StatusOK).SetMessage("Profile updated successfully").Build()
}
