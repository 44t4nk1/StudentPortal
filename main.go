package main

import (
	"log"
	"os"

	"github.com/44t4nk1/StudentPortal/api/routes"

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

	routes.InitRoutes(router)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}
