package main

import (
	"log"
	"os"
	"time"

	"github.com/44t4nk1/StudentPortal/api/middleware"
	"github.com/44t4nk1/StudentPortal/api/routes"
	cors "github.com/itsjamie/gin-cors"

	"github.com/44t4nk1/StudentPortal/api/db"
	"github.com/gin-gonic/gin"
)

var (
	err    error
	router = gin.Default()
)

func init() {
	_, err = db.RunDB()
	if err != nil {
		log.Fatal(err)
	}

	rateLimiter := middleware.RateLimitMiddleware()

	router.Use(
		cors.Middleware(
			cors.Config{
				Origins:         "*",
				Methods:         "GET, PUT, POST, DELETE",
				RequestHeaders:  "Origin, Authorization, Content-Type",
				ExposedHeaders:  "",
				MaxAge:          50 * time.Second,
				Credentials:     true,
				ValidateHeaders: false,
			}))
	router.Use(rateLimiter)
	routes.InitRoutes(router)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}
