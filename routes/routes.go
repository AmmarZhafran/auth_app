package routes

import (
	"auth-app/controllers"
	"auth-app/database"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	db := database.DB

	// Authentication routes
	r.POST("/register", controllers.Register(db))
	r.POST("/login", controllers.Login(db))
	r.POST("/verify-otp", controllers.VerifyOTP(db))

	return r
}
