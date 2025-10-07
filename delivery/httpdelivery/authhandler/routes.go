package authhandler

func (auth AuthHandler) AddRoutes() {
	routeGroup := auth.e.Group("/authenticate")

	routeGroup.POST("/auth", auth.Authenticate)
	routeGroup.POST("/verify", auth.Verify)
}
