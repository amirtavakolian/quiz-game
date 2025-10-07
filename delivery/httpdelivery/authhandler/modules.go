package authhandler

import (
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/notifier/sms"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository"
	"github.com/amirtavakolian/quiz-game/repository/otprepo"
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
)