package userservice

import (
	"context"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/param"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/helpers"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/logger"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/notifier/sms"
	responser "github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/responser"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/repository/otprepo"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/repository/userrepo"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/validator"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserService struct {
	UserValidator  validator.UserValidator
	UserRepository userrepo.UserRepositoryContract
	Responser      responser.Response
	Notifier       sms.SMSNotifier
	OTPRepository  otprepo.OTPContract
	Logger         logger.Logger
}

func NewUserService(userValidator validator.UserValidator, userRepo userrepo.UserRepositoryContract, response responser.Response, smsNotifier sms.SMSNotifier) UserService {
	return UserService{
		UserValidator:  userValidator,
		UserRepository: userRepo,
		Responser:      response,
		Notifier:       smsNotifier,
		OTPRepository:  otprepo.NewRedisOTPRepo(),
		Logger:         logger.New(),
	}
}

func (s UserService) Register(userParam param.RegisterParam) responser.Response {
	if status, errorsList := s.UserValidator.RegisterUserValidate(userParam); !status {
		return s.Responser.SetData(errorsList).SetStatusCode(http.StatusUnprocessableEntity).SetIsSuccess(false)
	}

	isUserRegisteredBefore, err := s.UserRepository.FindByPhoneNumber(userParam.PhoneNumber)

	if err != nil {
		//todo => add logger here
		return s.Responser.SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).SetIsSuccess(false)
	}

	if isUserRegisteredBefore {
		return s.Responser.SetMessage("You have registered before").SetStatusCode(http.StatusUnprocessableEntity).SetIsSuccess(false)
	}

	otpSixDigitCode, otpSixDigitCodeError := helpers.GenerateSixDigitCode()

	if otpSixDigitCodeError != nil {
		loggerSvc := s.Logger.Log()
		defer loggerSvc.Sync()
		loggerSvc.Error("Failed to generate OTP", zap.Error(err))
		return s.Responser.SetIsSuccess(false).SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).Build()
	}

	ctx := context.Background()
	ttl := 2 * time.Minute

	err = s.OTPRepository.Set(ctx, userParam.PhoneNumber+"_otp_code", otpSixDigitCode, ttl)

	if err != nil {
		loggerSvc := s.Logger.Log()
		defer loggerSvc.Sync()
		loggerSvc.Info("Redis set() error", zap.Error(err), zap.String("error message", err.Error()))
		return s.Responser.SetIsSuccess(false).SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).Build()
	}

	smsMessage := sms.NewSmsMessage()
	smsMessage.SetReceiverNumber(userParam.PhoneNumber).BuildCustomContent(sms.RegisterTemplate, otpSixDigitCode)
	sendSMSError := s.Notifier.SendSMS(smsMessage)

	if sendSMSError != nil {
		logger := s.Logger.Log()
		defer logger.Sync()
		return s.Responser.SetIsSuccess(true).SetMessage(sendSMSError.Error()).SetStatusCode(http.StatusInternalServerError).SetData(map[string]int{"ttl": int(ttl.Seconds())}).Build()

	}

	return s.Responser.SetIsSuccess(true).SetMessage("6 digit code sent to you").SetStatusCode(http.StatusOK).SetData(map[string]int{"ttl": int(ttl.Seconds())}).Build()
}
