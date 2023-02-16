package config

import (
	"github.com/spf13/viper"
)

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
