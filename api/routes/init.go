package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	router.GET("/home", GetFaculty)
}
