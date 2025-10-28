package main

import (
	"log"

	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery"
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/authhandler"
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/profilehandler"
	"github.com/amirtavakolian/quiz-game/pkg/jwt"
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/notifier/sms"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository"
	"github.com/amirtavakolian/quiz-game/repository/gorm/gormplayerrepo"
	"github.com/amirtavakolian/quiz-game/repository/gorm/gormprofilerepo"
	"github.com/amirtavakolian/quiz-game/repository/mysql/mysqlprofilerepo"
	"github.com/amirtavakolian/quiz-game/repository/otprepo"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	"github.com/amirtavakolian/quiz-game/service/authservice"
	"github.com/amirtavakolian/quiz-game/service/profileservice"
	"github.com/amirtavakolian/quiz-game/validator/auth"
	profilevalidator "github.com/amirtavakolian/quiz-game/validator/profile"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Modules = fx.Module(
	"serve",
	fx.Provide(repository.NewMysqlConnection),
	fx.Provide(repository.NewGormConnection),
	fx.Provide(auth.NewAuthValidator),
	fx.Provide(responser.NewResponse),
	fx.Provide(profilehandler.NewProfileHandler),
	fx.Provide(sms.NewNotifier),
	fx.Provide(otprepo.NewRedisOTPRepo),
	fx.Provide(logger.New),
	fx.Provide(authhandler.NewAuthHandler),
	fx.Provide(authservice.NewAuthService),
	fx.Provide(jwt.NewJwtService),
	fx.Provide(func() *echo.Echo { return echo.New() }),
	fx.Provide(httpdelivery.NewServe),
	fx.Provide(fx.Annotate(gormplayerrepo.NewPlayerRepo, fx.As(new(repositorycontracts.PlayerRepoContract)))),
	fx.Provide(fx.Annotate(gormprofilerepo.NewProfileRepo, fx.As(new(repositorycontracts.ProfileRepoContract)))),
	fx.Provide(profilevalidator.NewProfileValidator),
	fx.Provide(mysqlprofilerepo.NewProfileRepo),
	fx.Provide(profileservice.NewProfileService),
	fx.Invoke(func(s httpdelivery.Serve) {
		s.Serve()
	}),
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fx.New(
		Modules,
	).Run()
}
