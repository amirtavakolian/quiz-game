package httpdelivery

import (
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/authhandler"
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/profilehandler"
	"github.com/labstack/echo/v4"
)

type Serve struct {
	e          *echo.Echo
	authHld    authhandler.AuthHandler
	profileHld profilehandler.ProfileHandler
}

func NewServe(e *echo.Echo, authHld authhandler.AuthHandler, profileHld profilehandler.ProfileHandler) Serve {
	return Serve{
		e:          e,
		authHld:    authHld,
		profileHld: profileHld,
	}
}

func (s Serve) Serve() {
	s.loadRoutes()
	s.e.Logger.Fatal(s.e.Start(":80"))
}
