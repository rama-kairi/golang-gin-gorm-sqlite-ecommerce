package config

import (
	"github.com/spf13/viper"
)

// func InitConfig() {
// 	// load .env file
// 	err := godotenv.Load(".env.dev")
// 	if err != nil {
// 		panic("Error loading .env file")
// 	}
// }

func InitViperConfig() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env.dev")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Error loading .env file")
	}
}
