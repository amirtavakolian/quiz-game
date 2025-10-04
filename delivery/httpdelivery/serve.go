package httpdelivery

import (
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/authhandler"
	"github.com/labstack/echo/v4"
)

type Serve struct {
}

func (s Serve) Serve() {
	e := echo.New()
	LoadRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func LoadRoutes(e *echo.Echo) {
	authHld := authhandler.NewAuthHandler(e)
	authHld.AddRoutes()
}
