package main

import (
	"log"
	"os"

	"github.com/Danielyilma/Food_Recipes/services/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8090"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
