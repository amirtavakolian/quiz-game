package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
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
