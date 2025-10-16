package authservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/amirtavakolian/quiz-game/param/authparams"
	"github.com/amirtavakolian/quiz-game/pkg/helpers"
	"github.com/amirtavakolian/quiz-game/pkg/jwt"
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/notifier/sms"
	responser "github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository/otprepo"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	"github.com/amirtavakolian/quiz-game/validator/auth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Authenticate struct {
	Validator     auth.AuthValidator
	Responser     responser.Response
	Notifier      sms.SMSNotifier
	OTPRepository otprepo.OTPRepoContract
	Logger        logger.Logger
	PlayerRepo    repositorycontracts.PlayerRepoContract
	JWTService    *jwt.JWTService
}

const (
	OTPFailedAttempts   = "otp:fail:attempts:"
	OTPGeneratedCodeKey = "otp:code:"
	MaxWrongOTPAttempt  = "5"
)

func NewAuthService(
	validator auth.AuthValidator,
	responser responser.Response,
	notifier sms.SMSNotifier,
	otpRepo otprepo.RedisOTPRepo,
	logger logger.Logger,
	jwtService *jwt.JWTService,
	playerRepo repositorycontracts.PlayerRepoContract,
) Authenticate {
	return Authenticate{
		Validator:     validator,
		Responser:     responser,
		Notifier:      notifier,
		OTPRepository: otpRepo,
		Logger:        logger,
		JWTService:    jwtService,
		PlayerRepo:    playerRepo,
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
	/*
		smsMessage := sms.NewSmsMessage()
		smsMessage.SetReceiverNumber(userParam.PhoneNumber).BuildCustomContent(sms.RegisterTemplate, otpSixDigitCode)
		sendSMSError := s.Notifier.SendSMS(smsMessage)

		if sendSMSError != nil {
			loggerSvc.Error("Failed to generate OTP", zap.Error(sendSMSError), zap.String("error", sendSMSError.Error()))
			return s.Responser.SetIsSuccess(false).SetMessage("Internal server error").SetStatusCode(http.StatusInternalServerError).Build()
		}
	*/
	return s.Responser.SetIsSuccess(true).SetMessage("6 digit code sent to you").SetStatusCode(http.StatusOK).SetData(map[string]int{"ttl": int(ttl.Seconds())}).Build()
}

func (s Authenticate) Verify(verifyParam authparams.VerifyParam) responser.Response {
	ctx := context.Background()

	if err := s.Validator.Verify(verifyParam); err != nil {
		return s.Responser.SetData(err.Error()).SetStatusCode(400).Build()
	}

	OTPFailedAttemptsKey := OTPFailedAttempts + verifyParam.PhoneNumber
	err := s.checkOTPAttemptsLimit(ctx, OTPFailedAttemptsKey)

	if err != nil {
		return s.Responser.SetMessage(err.Error()).SetStatusCode(400).Build()
	}

	generatedOtpCode, err := s.OTPRepository.Get(ctx, OTPGeneratedCodeKey+verifyParam.PhoneNumber)
	if err != nil {
		if errors.Is(err, redis.Nil) { // key not exist
			return s.Responser.SetMessage("Please request for otp code first").SetStatusCode(400).Build()
		}
		return s.Responser.SetMessage("internal server error").SetStatusCode(500).Build()
	}

	ttl, err := s.validateOTPAndTrackAttempts(ctx, generatedOtpCode, verifyParam.OTPCode, OTPFailedAttemptsKey)
	if err != nil {
		response := s.Responser.SetMessage(err.Error()).SetStatusCode(400)

		if ttl != "" {
			response = response.SetData(map[string]string{"ttl": ttl})
		}
		return response.Build()
	}

	token, err := s.JWTService.GenerateToken(verifyParam.PhoneNumber)
	if err != nil {
		s.Logger.Log().Error("generate token", zap.Error(err), zap.String("generate-token", err.Error()))
		return s.Responser.SetMessage("internal server error").SetStatusCode(500).Build()
	}

	if _, err = s.OTPRepository.Del(ctx, OTPFailedAttemptsKey, OTPGeneratedCodeKey+verifyParam.PhoneNumber); err != nil {
		s.Logger.Log().Error("delete key from redis", zap.Error(err), zap.String("delete-key", err.Error()))
		return s.Responser.SetMessage("internal server error").SetStatusCode(500).Build()
	}

	if err = s.PlayerRepo.Store(verifyParam.PhoneNumber); err != nil {
		s.Logger.Log().Error("database error", zap.Error(err), zap.String("insert-phone", err.Error()))
		return s.Responser.SetMessage("internal server error").SetStatusCode(500).Build()
	}

	return s.Responser.SetStatusCode(200).SetIsSuccess(true).SetData(map[string]string{"token": token}).Build()
}

func (s Authenticate) checkOTPAttemptsLimit(ctx context.Context, OTPFailedAttemptsKey string) error {
	OTPFailedAttemptsCount, err := s.OTPRepository.Get(ctx, OTPFailedAttemptsKey)

	if err != redis.Nil {
		if OTPFailedAttemptsCount == MaxWrongOTPAttempt {
			ttl, err := s.OTPRepository.TTL(ctx, OTPFailedAttemptsKey)

			if err != nil {
				s.Logger.Log().Error("redis ttl lookup failed", zap.Error(err))
				return errors.New("internal server error")
			}

			msg := fmt.Sprintf("You have entered the code incorrectly 5 times. Please try again after %s", ttl.String())
			return errors.New(msg)
		}
	}
	return nil
}

func (s Authenticate) validateOTPAndTrackAttempts(ctx context.Context, generatedOtpCode string, userOtp string, OTPFailedAttemptsKey string) (string, error) {
	if generatedOtpCode != userOtp {
		count, err := s.OTPRepository.Incr(ctx, OTPFailedAttemptsKey)

		if err != nil {
			return "", err
		}

		if strconv.Itoa(int(count)) == MaxWrongOTPAttempt {
			_, err := s.OTPRepository.Expire(ctx, OTPFailedAttemptsKey, 5*time.Minute)

			if err != nil {
				return "", err
			}

			ttl, err := s.OTPRepository.TTL(ctx, OTPFailedAttemptsKey)

			msg := fmt.Sprintf("you have entered the code incorrectly 5 times. Please try again after %s.")
			return ttl.String(), errors.New(msg)
		}

		return "", errors.New("the code you entered is incorrect")
	}
	return "", nil
}
