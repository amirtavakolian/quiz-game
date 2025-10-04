package authhandler

import (
	"github.com/amirtavakolian/quiz-game/param/authparams"
	"github.com/amirtavakolian/quiz-game/repository"
	"github.com/amirtavakolian/quiz-game/repository/mysql/playerrepo"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	"github.com/amirtavakolian/quiz-game/service/authservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	e *echo.Echo
}

func NewAuthHandler(e *echo.Echo) AuthHandler {
	return AuthHandler{e: e}
}

func (auth AuthHandler) Authenticate(c echo.Context) error {
	var params authparams.RegisterParam
	var playerrepository repositorycontracts.PlayerRepoContract

	playerrepository = playerrepo.NewPlayerRepo(repository.NewMysqlConnection())
	authSvc := authservice.NewAuthService(playerrepository)

	if err := c.Bind(&params); err != nil {
		// logg error
	}

	result := authSvc.Authenticate(params)
	return c.JSON(http.StatusOK, result)
}
