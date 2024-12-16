package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/gin-gonic/gin"
)

func Payment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Payload
		url := "https://api.chapa.co/v1/transaction/initialize"

		if err := c.BindJSON(&payload); err != nil {
			log.Println("Invalid JSON received:", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		// Convert the Payload struct into JSON
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("Error marshalling payload: %v", err)
		}

		// Create a new HTTP request with the given URL, method, and body
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		// Set headers for the request
		req.Header.Set("Authorization", "Bearer CHASECK_TEST-RZLz1c03oPMLdDMUCHOq8rZZKmgSHw5v")
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		defer resp.Body.Close()

		// Read and print the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Print the response text
		c.JSON(http.StatusOK, gin.H{"success": string(body)})
	}
}
