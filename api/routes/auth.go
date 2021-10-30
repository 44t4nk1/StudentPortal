package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFaculty(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Home Page",
	})
}
