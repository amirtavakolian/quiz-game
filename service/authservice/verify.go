package authservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/amirtavakolian/quiz-game/param/authparams"
	"github.com/amirtavakolian/quiz-game/pkg/jwt"
	"github.com/amirtavakolian/quiz-game/pkg/logger"
	"github.com/amirtavakolian/quiz-game/pkg/notifier/sms"
	"github.com/amirtavakolian/quiz-game/pkg/responser"
	"github.com/amirtavakolian/quiz-game/repository/otprepo"
	"github.com/amirtavakolian/quiz-game/repository/repositorycontracts"
	"github.com/amirtavakolian/quiz-game/validator/auth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
	"time"
)

const (
	OTPFailedAttempts   = "otp:fail:attempts:"
	OTPGeneratedCodeKey = "otp:code:"
	MaxWrongOTPAttempt  = "5"
)

type Verify struct {
	Validator     auth.AuthValidator
	playerRepo    repositorycontracts.PlayerRepoContract
	Responser     responser.Response
	Notifier      sms.SMSNotifier
	OTPRepository otprepo.OTPRepoContract
	Logger        logger.Logger
	JWTService    *jwt.JWTService
}

func NewVerifyService(playerRepo repositorycontracts.PlayerRepoContract, jwt *jwt.JWTService) Verify {
	return Verify{
		Validator:     auth.NewAuthValidator(),
		playerRepo:    playerRepo,
		Responser:     responser.NewResponse(),
		Notifier:      sms.NewNotifier(),
		OTPRepository: otprepo.NewRedisOTPRepo(),
		Logger:        logger.New(),
		JWTService:    jwt,
	}
}

func (s Verify) Verify(verifyParam *authparams.VerifyParam) responser.Response {
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

	token, err := s.JWTService.GenerateToken()
	if err != nil {
		s.Logger.Log().Error("generate token", zap.Error(err), zap.String("generate-token", err.Error()))
		return s.Responser.SetMessage("internal server error").SetStatusCode(500).Build()
	}

	_, err = s.OTPRepository.Del(ctx, OTPFailedAttemptsKey, OTPGeneratedCodeKey+verifyParam.PhoneNumber)
	if err != nil {
		s.Logger.Log().Error("delete key from redis", zap.Error(err), zap.String("delete-key", err.Error()))
		return s.Responser.SetMessage("internal server error").SetStatusCode(500).Build()
	}

	return s.Responser.SetStatusCode(200).SetIsSuccess(true).SetData(map[string]string{"token": token}).Build()
}

func (s Verify) checkOTPAttemptsLimit(ctx context.Context, OTPFailedAttemptsKey string) error {
	OTPFailedAttemptsCount, err := s.OTPRepository.Get(ctx, OTPFailedAttemptsKey)

	if err != redis.Nil {
		if OTPFailedAttemptsCount == MaxWrongOTPAttempt {
			ttl, err := s.OTPRepository.TTL(ctx, OTPFailedAttemptsKey)

			if err != nil {
				s.Logger.Log().Error("redis ttl lookup failed", zap.Error(err))
				return errors.New("internal server error")
			}

			msg := fmt.Sprintf("You have entered the code incorrectly 5 times. Please try again after %s.", ttl.String())
			return errors.New(msg)
		}
	}
	return nil
}

func (s Verify) validateOTPAndTrackAttempts(ctx context.Context, generatedOtpCode string, userOtp string, OTPFailedAttemptsKey string) (string, error) {
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
