package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/middleware"
)

func productRoutes(e *gin.Engine, client *ent.Client) {
	productApi := controllers.NewProductController(client)

	productGroup := e.Group("/product")
	{
		productGroup.GET("", productApi.GetAll)
		productGroup.POST("", productApi.Create, middleware.AuthMiddleware(client))
		productGroup.GET("/:id", productApi.Get)
		productGroup.DELETE("/:id", productApi.Delete)
		productGroup.PATCH("/:id", productApi.Update)
		productGroup.GET("/user/:id", productApi.GetAllByUser)
	}
}
