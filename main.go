package main

import (
	"os"

	"github.com/rama-kairi/blog-api-golang-gin/config"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/routes"
)

func main() {
	config.InitConfig()

	db.InitGormDb()

	// Auto migrate the models
	db.Db.AutoMigrate(&models.User{})

	r := routes.InitRoutes()

	r.Run(":" + os.Getenv("APP_PORT"))
}
