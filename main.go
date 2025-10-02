package main

import (
	"fmt"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/param"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/notifier/sms"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/responser"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/repository/userrepo"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/service/authservice"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/validator/auth"

)

func main() {
	userValidator := auth.NewUserValidator()
	inMemoryRepo := userrepo.NewInMemoryUserRepo()
	responseSvc := responser.NewResponse()
	smsNotifi := sms.NewNotifier()
	userSvc := authservice.NewUserService(userValidator, inMemoryRepo, responseSvc, smsNotifi)
	userInput := GetUserInput()

	result := userSvc.Authenticate(userInput)

	fmt.Print(result)
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
