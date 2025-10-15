package httpdelivery

import (
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/authhandler"
	"github.com/labstack/echo/v4"
)

type Serve struct {
	e       *echo.Echo
	authHld authhandler.AuthHandler
}

func NewServe(e *echo.Echo, authHld authhandler.AuthHandler) Serve {
	return Serve{
		e:       e,
		authHld: authHld,
	}
}

func (s Serve) Serve() {
	s.loadRoutes()
	s.e.Logger.Fatal(s.e.Start(":80"))
}

func (s Serve) loadRoutes() {
	routeGroup := s.e.Group("/authenticate")
	routeGroup.POST("/auth", s.authHld.Authenticate)
	routeGroup.POST("/verify", s.authHld.Verify)
}
