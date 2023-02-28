package routes

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()

	userRoutes(r)        // routes/user.go
	authRoutes(r)        // routes/auth.go
	productRoutes(r)     // routes/product.go
	categoryRoutes(r)    // routes/category.go
	subCategoryRoutes(r) // routes/sub_category.go

	return r
}
