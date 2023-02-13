package main

import (
	"github.com/rama-kairi/blog-api-golang-gin/config"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/routes"
	"github.com/spf13/viper"
)

func main() {
	config.InitViperConfig()

	db.InitEntDb()

	r := routes.InitRoutes()

	r.Run(":" + viper.GetString("APP_PORT"))
}
