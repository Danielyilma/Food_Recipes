package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Danielyilma/Food_Recipes/services/helpers"
	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/Danielyilma/Food_Recipes/services/query"
	"github.com/gin-gonic/gin"
)

func UpdateRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipe models.Recipe

		id := c.DefaultPostForm("recipe_id", "")
		oldRecipe, err := query.GetRecipe(id, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recipe.Id = oldRecipe.Id
		// Parse basic fields
		recipe.Title = c.DefaultPostForm("title", "")
		recipe.Description = c.DefaultPostForm("description", "")
		recipe.Thumbnail = c.DefaultPostForm("featuredimage", "")
		helpers.DeleteFile(oldRecipe.Thumbnail)
		prepTime, err := strconv.Atoi(c.DefaultPostForm("prepTime", ""))

		if err != nil {
			log.Printf("Error converting category_id to integer: %v", err)
			prepTime = 30
		}
		recipe.PrepTime = prepTime

		categoryID, err := strconv.Atoi(c.DefaultPostForm("category_id", ""))

		if err != nil {
			log.Printf("Error converting prepTime to integer: %v", err)
			categoryID = 0
		}
		recipe.CategoryID = categoryID

		// Parse ingredients
		ingredientCount := getFormArrayCount(c, "ingredient", ".name")
		var ingredients []models.Recipe_ingredients
		for i := 0; i < ingredientCount; i++ {
			name := c.DefaultPostForm(fmt.Sprintf("ingredient[%d].name", i), "")
			amount := c.DefaultPostForm(fmt.Sprintf("ingredient[%d].amount", i), "")
			unit := c.DefaultPostForm(fmt.Sprintf("ingredient[%d].unit", i), "")
			ingredients = append(ingredients, models.Recipe_ingredients{Ingredient: models.Ingredient{Name: name}, Amount: amount, Unit: unit})
		}
		recipe.Ingredients = ingredients

		// Parse steps
		stepCount := getFormArrayCount(c, "steps", ".description")
		var steps []models.Step
		for i := 0; i < stepCount; i++ {
			Instruction := c.DefaultPostForm(fmt.Sprintf("steps[%d].description", i), "")
			image_path := c.DefaultPostForm(fmt.Sprintf("steps[%d].image", i), "")
			helpers.DeleteFile(oldRecipe.Steps[i].RecipeImages.ImageURL)
			step_number, _ := strconv.Atoi(c.DefaultPostForm(fmt.Sprintf("steps[%d].step_number", i), ""))
			// path, err := helpers.UploadFile(image)
			log.Print(image_path)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			steps = append(steps, models.Step{StepNumber: step_number, Instruction: Instruction, RecipeImages: models.RecipeImage{ImageURL: image_path}})
		}
		recipe.Steps = steps

		imageCount := getFormArrayCount(c, "images", "")
		var images []models.RecipeImage

		for i := 0; i < imageCount; i++ {
			image_path := c.DefaultPostForm(fmt.Sprintf("images[%d]", i), "")

			// helpers.DeleteFile(oldRecipe.Images[i].ImageURL)
			// log.Print("eeeeeeeeeeeeeeee")
			log.Print(image_path)
			// path, err := helpers.UploadFile(image)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if i == 0 {
				recipe.Thumbnail = image_path
			}
			images = append(images, models.RecipeImage{ImageURL: image_path})
		}
		recipe.Images = images
		response, err := query.SendUpdateMutation(query.UpdateMutation, recipe, c)

		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": *response})
	}
}
