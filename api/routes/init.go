package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)
}
