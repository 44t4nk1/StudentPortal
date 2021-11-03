package routes

import (
	"net/http"

	"github.com/44t4nk1/StudentPortal/api/db"
	"github.com/44t4nk1/StudentPortal/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetStudentData(c *gin.Context, token *jwt.Token) {
	var student models.Student

	id, err := uuid.Parse(token.Claims.(jwt.MapClaims)["id"].(string))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid JWT Provided"})
		return
	}
	DB := db.GetDB()
	err = DB.Debug().Model(models.Student{}).Where("uuid = ?", id).Take(&student).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error in DB"})
		return
	}
	var studentData models.StudentData

	studentData.UUID = student.UUID
	studentData.Email = student.Email
	studentData.Name = student.Name
	studentData.RegNo = student.RegNo
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": studentData})
}
