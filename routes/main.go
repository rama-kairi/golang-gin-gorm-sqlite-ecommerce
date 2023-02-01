package routes

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()

	userRoutes(r) // routes/user.go
	authRoutes(r) // routes/auth.go

	return r
}
