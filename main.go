package main

import (
	"fmt"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/param"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/configloader"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/jwt"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/logger"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/notifier/sms"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/responser"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/repository/otprepo"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/repository/userrepo"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/service/authservice"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/validator/auth"
)

func main() {
	// authenticate()
	verify()
}

func GetUserInput() param.RegisterParam {
	var registerParam param.RegisterParam

	fmt.Print("\n Enter your first name: ")
	fmt.Scan(&registerParam.Name)

	fmt.Print("\nEnter your family: ")
	fmt.Scan(&registerParam.Family)

	fmt.Print("\nEnter your phone number: ")
	fmt.Scan(&registerParam.PhoneNumber)

	return registerParam
}

func authenticate() {
	userValidator := auth.NewUserValidator()
	inMemoryRepo := userrepo.NewInMemoryUserRepo()
	responseSvc := responser.NewResponse()
	smsNotifi := sms.NewNotifier()
	userSvc := authservice.NewUserService(userValidator, inMemoryRepo, responseSvc, smsNotifi)
	userInput := GetUserInput()

	result := userSvc.Authenticate(userInput)

	fmt.Print(result)
}

func verify() {
	verifyParams := param.VerifyParam{
		PhoneNumber: "09120000000",
		OTPCode:     "585715",
	}

	cfgLoader := configloader.NewConfigLoader()
	result := cfgLoader.SetPrefix("APP_").SetDelimiter(".").SetDivider("_").Build()

	secretKey := []byte(result.String("jwt.secret.key"))

	verifySvc := authservice.Verify{
		Validator:     auth.NewVerifyValidator(),
		Responser:     responser.NewResponse(),
		OTPRepository: otprepo.NewRedisOTPRepo(),
		JWTService:    jwt.NewJwtService(secretKey),
		LoggerSvc:     logger.New(),
	}

	result12 := verifySvc.Verify(verifyParams)

	fmt.Println(result12)
}
