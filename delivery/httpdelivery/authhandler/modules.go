package authhandler

import (
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
	"go.uber.org/fx"
)

var Modules = fx.Module(
	"authhandler",
	fx.Provide(repository.NewMysqlConnection),
	fx.Provide(auth.NewAuthValidator),
	fx.Provide(responser.NewResponse),
	fx.Provide(sms.NewNotifier),
	fx.Provide(otprepo.NewRedisOTPRepo),
	fx.Provide(logger.New),
	fx.Provide(authservice.NewAuthService),
	fx.Provide(func() []byte {
		cfgLoader := configloader.NewConfigLoader()
		result := cfgLoader.SetPrefix("APP_").SetDelimiter(".").SetDivider("_").Build()
		return []byte(result.String("jwt.secret.key"))
	}),
	fx.Provide(jwt.NewJwtService),
	fx.Provide(fx.Annotate(
		playerrepo.NewPlayerRepo,
		fx.As(new(repositorycontracts.PlayerRepoContract)),
	)),
)
