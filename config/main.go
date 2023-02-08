package config

import (
	"github.com/joho/godotenv"
)

var configPath = ".env.dev"

func InitConfig() {
	// load .env file
	err := godotenv.Load(configPath)
	if err != nil {
		panic("Error loading .env file")
	}
}
