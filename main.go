package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
)

func main() {
	db.InitGormDb()

	// Auto migrate the models
	db.Db.AutoMigrate(&models.User{})

	r := gin.Default()
	userApi := controllers.NewUserController()

	r.GET("/user", userApi.GetAll)
	r.POST("/user", userApi.Create)
	r.GET("/user/:id", userApi.Get)
	r.DELETE("/user/:id", userApi.Delete)
	r.PATCH("/user/:id", userApi.Update)

	r.Run(":8080")
}
