package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/Danielyilma/Food_Recipes/services/query"
	"github.com/gin-gonic/gin"
)

func getFormArrayCount(c *gin.Context, name string, field string) int {
	count := 0
	for {
		_, exists := c.GetPostForm(fmt.Sprintf("%s[%d]%s", name, count, field))
		if !exists {
			break
		}
		count++
	}
	return count
}

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipe models.Recipe

		// Parse basic fields
		recipe.Title = c.DefaultPostForm("title", "")
		recipe.Description = c.DefaultPostForm("description", "")
		recipe.Thumbnail = c.DefaultPostForm("featuredimage", "")
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

		response, err := query.SendMutation(recipe, c)

		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": *response})
	}
}

// mutation InsertRecipe {
// 	insert_recipes_one(
// 	  object: {
// 		title: "Delicious Pancakes"
// 		description: "Fluffy pancakes made with love."
// 		prep_time: 30
// 		thumbnail: "https://example.com/images/pancakes.jpg"
// 		is_paid: true
// 		price: 6.00
// 		user_id: 1
// 		category_id: 2
// 		recipe_ingredients: {
// 		  data: [
// 			{
// 			  ingredient: {
// 				data: { name: "Flour", description: "All-purpose flour" }
// 			  }
// 			  quantity: "2 cups"
// 			}
// 			{
// 			  ingredient: {
// 				data: { name: "Sugar", description: "Granulated sugar" }
// 			  }
// 			  quantity: "1 tbsp"
// 			}
// 			{
// 			  ingredient: {
// 				data: { name: "Milk", description: "Whole milk" }
// 			  }
// 			  quantity: "1 cup"
// 			}
// 			{
// 			  ingredient: {
// 				data: { name: "Egg", description: "Large egg" }
// 			  }
// 			  quantity: "2"
// 			}
// 		  ]
// 		}
// 		steps: {
// 		  data: [
// 			{ step_number: 1, instruction: "Mix all dry ingredients." , recipe_images: {data: {image_url: "https://example.com/images/step1.jpg"}}}
// 			{ step_number: 2, instruction: "Add milk and eggs, then stir well.", recipe_images: {data: {image_url: "https://example.com/images/step1.jpg"}}}
// 			{ step_number: 3, instruction: "Cook on a hot griddle until golden brown.", recipe_images: {data: {image_url: "https://example.com/images/step1.jpg"}}}
// 		  ]
// 		}
// 		recipe_images: {
// 		  data: [
// 			{ image_url: "https://example.com/images/step1.jpg" }
// 			{ image_url: "https://example.com/images/step2.jpg" }
// 		  ]
// 		}
// 	  }
// 	) {
// 	  id
// 	  title
// 	  created_at
// 	}

// imgPath, err := helpers.UploadByName(c, "featuredimage", nil)

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"errort": err.Error()})
// 			return
// 		}
// 		recipe.Thumbnail = *imgPath

// 		form, err := c.MultipartForm()
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"errork": err.Error()})
// 			return
// 		}

// 		imageFiles := form.File["images"]

// 		for _, image := range imageFiles {
// 			if helpers.CheckFileExists(c, &image.Filename) {
// 				continue
// 			}

// 			imgPath, err := helpers.UploadByName(c, "", image)
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}
// 			recipeImage = append(recipeImage, models.RecipeImage{ImageURL: *imgPath})
// 		}

// 		c.JSON(http.StatusOK, gin.H{"recipe": recipe, "images": recipeImage})
