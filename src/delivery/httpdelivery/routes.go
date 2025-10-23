package httpdelivery

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"os"
)

func (s Serve) loadRoutes() {
	authRouteGroup := s.e.Group("/authenticate")
	authRouteGroup.POST("/auth", s.authHld.Authenticate)
	authRouteGroup.POST("/verify", s.authHld.Verify)

	profileRouteGroup := s.e.Group("/profile", echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	profileRouteGroup.POST("/update", s.profileHld.Update)
}
