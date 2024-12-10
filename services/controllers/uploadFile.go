package controllers

import (
	"net/http"

	"github.com/Danielyilma/Food_Recipes/services/helpers"
	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.DefaultPostForm("file", "")

		path, err := helpers.UploadFile(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"url":     path,
		})
	}
}
