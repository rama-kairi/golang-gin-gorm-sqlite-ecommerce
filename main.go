package main

import (
	"github.com/rama-kairi/blog-api-golang-gin/config"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/routes"
	"github.com/spf13/viper"
)

func main() {
	config.InitViperConfig()

	// Initialize the Ent ORM
	entClient := db.InitEntDb()
	defer entClient.Close() // Close the Ent ORM

	r := routes.InitRoutes(entClient)

	r.Run(":" + viper.GetString("APP_PORT"))
}
