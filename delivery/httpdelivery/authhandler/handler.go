package authhandler

import (
	"github.com/amirtavakolian/quiz-game/param/authparams"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/service/authservice"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
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
	var result responser.Response

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fx.New(
		Modules,
		fx.Invoke(func(s authservice.Authenticate) {
			result = s.Authenticate(params)
		}),
	)

	return c.JSON(http.StatusOK, result)
}
