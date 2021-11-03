package routes

import (
	"github.com/44t4nk1/StudentPortal/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)
	router.GET("/home", middleware.IsAuth(GetStudentData))
}
