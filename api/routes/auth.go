package routes

import (
	"log"
	"net/http"

	"github.com/44t4nk1/StudentPortal/api/db"
	"github.com/44t4nk1/StudentPortal/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func Signup(c *gin.Context) {
	var student models.StudentReg
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid JSON Provided"})
		return
	}
	DB := db.GetDB()
	var result models.Student
	err := DB.Debug().Model(models.Student{}).Where("email = ?", student.Email).Take(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		studentRecord := models.Student{
			UUID:     uuid.New(),
			Email:    student.Email,
			Name:     student.Name,
			RegNo:    student.RegNo,
			Password: hashAndSalt(([]byte(student.Password))),
		}
		err = DB.Debug().Create(&studentRecord).Error
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "Account created succesfully",
		})
	} else if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error in accessing DB"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
	}
}
