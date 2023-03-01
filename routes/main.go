package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
)

func InitRoutes(client *ent.Client) *gin.Engine {
	r := gin.Default()

	userRoutes(r, client)        // routes/user.go
	authRoutes(r, client)        // routes/auth.go
	productRoutes(r, client)     // routes/product.go
	categoryRoutes(r, client)    // routes/category.go
	subCategoryRoutes(r, client) // routes/sub_category.go

	return r
}
