package main

import (
	"log"
	"os"

	"github.com/Danielyilma/Food_Recipes/services/routes"
	"github.com/gin-contrib/cors"
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
	CORS := os.Getenv("CORS")
	if CORS == "" {
		CORS = "http://localhost:3000"
	}

	router := gin.New()
	router.MaxMultipartMemory = 50 * 1024 * 1024
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{CORS}, // Allow requests from your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Static("/uploads", "../uploads/")

	routes.AuthRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
