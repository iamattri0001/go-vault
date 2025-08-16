package main

import (
	"go-vault/bootstrap"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	app, err := bootstrap.NewApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err.Error())
	}

	defer app.Close()
	app.Run()
}
