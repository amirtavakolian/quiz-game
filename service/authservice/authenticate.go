package authservice

import (
	"context"
	"github.com/amirtavakolian/quiz-game/param/authparams"
	"github.com/amirtavakolian/quiz-game/pkg/helpers"
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/notifier/sms"
	responser "github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository/otprepo"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	"github.com/amirtavakolian/quiz-game/validator/auth"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Authenticate struct {
	Validator     auth.AuthValidator
	playerRepo    repositorycontracts.PlayerRepoContract
	Responser     responser.Response
	Notifier      sms.SMSNotifier
	OTPRepository otprepo.OTPRepoContract
	Logger        logger.Logger
}

func NewAuthService(playerRepo repositorycontracts.PlayerRepoContract) Authenticate {
	return Authenticate{
		Validator:     auth.NewAuthValidator(),
		playerRepo:    playerRepo,
		Responser:     responser.NewResponse(),
		Notifier:      sms.NewNotifier(),
		OTPRepository: otprepo.NewRedisOTPRepo(),
		Logger:        logger.New(),
	}
}

func (s Authenticate) Authenticate(userParam authparams.RegisterParam) responser.Response {
	if err := s.Validator.Authenticate(userParam); err != nil {
		return s.Responser.SetData(err.Error()).SetStatusCode(http.StatusUnprocessableEntity).SetIsSuccess(false)
	}

	otpSixDigitCode, otpSixDigitCodeError := helpers.GenerateSixDigitCode()

	loggerSvc := s.Logger.Log()
	defer loggerSvc.Sync()

	if otpSixDigitCodeError != nil {
		loggerSvc.Error("Failed to generate OTP", zap.Error(otpSixDigitCodeError), zap.String("error", otpSixDigitCodeError.Error()))
		return s.Responser.SetIsSuccess(false).SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).Build()
	}

	ctx := context.Background()
	ttl := 10 * time.Minute
	err := s.OTPRepository.Set(ctx, OTPGeneratedCodeKey+userParam.PhoneNumber, otpSixDigitCode, ttl)

	if err != nil {
		loggerSvc.Info("Redis set() error", zap.Error(err), zap.String("error message", err.Error()))
		return s.Responser.SetIsSuccess(false).SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).Build()
	}

	smsMessage := sms.NewSmsMessage()
	smsMessage.SetReceiverNumber(userParam.PhoneNumber).BuildCustomContent(sms.RegisterTemplate, otpSixDigitCode)
	sendSMSError := s.Notifier.SendSMS(smsMessage)

	if sendSMSError != nil {
		loggerSvc.Error("Failed to generate OTP", zap.Error(sendSMSError), zap.String("error", sendSMSError.Error()))
		return s.Responser.SetIsSuccess(false).SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).Build()
	}

	return s.Responser.SetIsSuccess(true).SetMessage("6 digit code sent to you").SetStatusCode(http.StatusOK).SetData(map[string]int{"ttl": int(ttl.Seconds())}).Build()
}
