package query

import (
	"context"
	"fmt"
	"log"
	"os"

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

func SendMutation(mutation string, recipe models.Recipe, c *gin.Context) (*models.InsertRecipesOne, error) {
	GRAPHQL_API := os.Getenv("GRAPHQL_API")
	if GRAPHQL_API == "" {
		GRAPHQL_API = "http://localhost:8080/v1/graphql"
	}
	client := graphql.NewClient(GRAPHQL_API)

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

	if recipe.Id != nil {
		req.Var("recipe_id", *recipe.Id)
	}

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
