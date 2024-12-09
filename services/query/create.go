package query

import (
	"context"
	"fmt"
	"log"

	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

func mapRecipeImages(images []models.RecipeImage) map[string]interface{} {
	return map[string]interface{}{
		"data": func() []map[string]interface{} {
			var result []map[string]interface{}
			for _, image := range images {
				result = append(result, map[string]interface{}{
					"image_url": image.ImageURL,
				})
			}
			return result
		}(),
	}
}

func mapRecipeSteps(steps []models.Step) map[string]interface{} {
	return map[string]interface{}{
		"data": func() []map[string]interface{} {
			var result []map[string]interface{}
			for _, step := range steps {
				result = append(result, map[string]interface{}{
					"step_number": step.StepNumber,
					"instruction": step.Instruction,
					"recipe_images": map[string]interface{}{
						"data": []map[string]interface{}{
							{
								"image_url": step.RecipeImages.ImageURL,
							},
						},
					},
				})
			}
			return result
		}(),
	}
}

func SendMutation(recipe models.Recipe, c *gin.Context) (*models.InsertRecipesOne, error) {
	client := graphql.NewClient("http://localhost:8080/v1/graphql")

	// Define the mutation
	mutation := `
	mutation InsertRecipe($title: String!, $description: String!, $prepTime: Int!, $thumbnail: String, $isPaid: Boolean, $price: numeric, $categoryId: Int!, $ingredients: recipe_ingredients_arr_rel_insert_input, $steps: steps_arr_rel_insert_input, $images: recipe_images_arr_rel_insert_input) {
		insert_recipes_one(object: {
			title: $title
			description: $description
			prep_time: $prepTime
			thumbnail: $thumbnail
			is_paid: $isPaid
			price: $price
			category_id: $categoryId
			recipe_ingredients: $ingredients
			steps: $steps
			recipe_images: $images
		}) {
			id
			title
			created_at
		}
	}`

	// Convert the struct into a format that GraphQL can accept
	ingredients := map[string]interface{}{
		"data": make([]map[string]interface{}, len(recipe.Ingredients)),
	}

	for i, ri := range recipe.Ingredients {
		ingredients["data"].([]map[string]interface{})[i] = map[string]interface{}{
			"ingredient": map[string]interface{}{
				"data": map[string]interface{}{
					"name": ri.Ingredient.Name,
				},
			},
			"quantity": ri.Amount,
			"unit":     ri.Unit, // Amount as quantity
		}
	}
	log.Print(ingredients)

	steps := mapRecipeSteps(recipe.Steps)

	images := mapRecipeImages(recipe.Images)

	// Create a request
	req := graphql.NewRequest(mutation)
	req.Header.Set("Authorization", c.GetHeader("Authorization"))

	// Set the variables dynamically from the recipe struct
	req.Var("title", recipe.Title)
	req.Var("description", recipe.Description)
	req.Var("prepTime", recipe.PrepTime)
	req.Var("thumbnail", recipe.Thumbnail)
	req.Var("isPaid", recipe.IsPaid)
	req.Var("price", fmt.Sprintf("%.2f", recipe.Price))
	req.Var("categoryId", recipe.CategoryID)
	req.Var("ingredients", ingredients)
	req.Var("steps", steps)
	req.Var("images", images)

	// Optional: Add authentication header
	// req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

	// Execute the request
	var response struct {
		InsertRecipesOne models.InsertRecipesOne `json:"insert_recipes_one"`
	}

	err := client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	}

	// Output the response
	// fmt.Printf("Recipe created: %+v\n", response.InsertRecipesOne)

	return (&response.InsertRecipesOne), nil
}
