package authhandler

import (
	"github.com/amirtavakolian/quiz-game/param/authparams"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/service/authservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	AuthService authservice.Authenticate
}

func NewAuthHandler(authService authservice.Authenticate) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

func (auth AuthHandler) Authenticate(c echo.Context) error {
	var params authparams.RegisterParam
	var result responser.Response

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result = auth.AuthService.Authenticate(params)
	return c.JSON(http.StatusOK, result)
}

func (auth AuthHandler) Verify(c echo.Context) error {
	var verifyParams authparams.VerifyParam
	var result responser.Response

	if err := c.Bind(&verifyParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result = auth.AuthService.Verify(verifyParams)
	return c.JSON(http.StatusOK, result)
}
