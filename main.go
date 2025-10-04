package main

import (
	"github.com/amirtavakolian/quiz-game/delivery/httpdelivery"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serve := httpdelivery.Serve{}
	serve.Serve()
}
