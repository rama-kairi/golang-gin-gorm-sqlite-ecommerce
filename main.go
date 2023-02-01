package main

import (
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/routes"
)

func main() {
	db.InitGormDb()

	// Auto migrate the models
	db.Db.AutoMigrate(&models.User{})

	r := routes.InitRoutes()
	r.Run(":8080")
}
