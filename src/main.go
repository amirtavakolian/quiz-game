package main

import (
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery"
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery/authhandler"
	"github.com/amirtavakolian/quiz-game/pkg/configloader"
	"github.com/amirtavakolian/quiz-game/pkg/jwt"
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/notifier/sms"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository"
	"github.com/amirtavakolian/quiz-game/repository/mysql/playerrepo"
	"github.com/amirtavakolian/quiz-game/repository/otprepo"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	"github.com/amirtavakolian/quiz-game/service/authservice"
	"github.com/amirtavakolian/quiz-game/validator/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"log"
)

var Modules = fx.Module(
	"serve",
	fx.Provide(repository.NewMysqlConnection),
	fx.Provide(auth.NewAuthValidator),
	fx.Provide(responser.NewResponse),
	fx.Provide(sms.NewNotifier),
	fx.Provide(otprepo.NewRedisOTPRepo),
	fx.Provide(logger.New),
	fx.Provide(authhandler.NewAuthHandler),
	fx.Provide(authservice.NewAuthService),
	fx.Provide(func() []byte {
		cfgLoader := configloader.NewConfigLoader()
		result := cfgLoader.SetPrefix("APP_").SetDelimiter(".").SetDivider("_").Build()
		return []byte(result.String("jwt.secret.key"))
	}),
	fx.Provide(jwt.NewJwtService),
	fx.Provide(func() *echo.Echo { return echo.New() }),
	fx.Provide(httpdelivery.NewServe),
	fx.Provide(fx.Annotate(
		playerrepo.NewPlayerRepo,
		fx.As(new(repositorycontracts.PlayerRepoContract)),
	)),
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
