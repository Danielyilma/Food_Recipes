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

var InsertMutation = `
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

var UpdateMutation = `
mutation UpdateRecipe(
  $recipe_id: Int!,
  $title: String,
  $description: String,
  $prepTime: Int,
  $thumbnail: String,
  $isPaid: Boolean,
  $price: numeric,
  $categoryId: Int,
  $ingredients: [recipe_ingredients_insert_input!]!,
  $steps: [steps_insert_input!]!,
  $images: [recipe_images_insert_input!]!
) {
  update_recipes_by_pk(
    pk_columns: { id: $recipe_id },
    _set: {
      title: $title,
      description: $description,
      prep_time: $prepTime,
      thumbnail: $thumbnail,
      is_paid: $isPaid,
      price: $price,
      category_id: $categoryId
    }
  ) {
    id
  }
  delete_recipe_ingredients(where: { recipe_id: { _eq: $recipe_id } }) {
    returning { id }
  }
  insert_recipe_ingredients(objects: $ingredients) {
    returning { id }
  }
  delete_steps(where: { recipe_id: { _eq: $recipe_id } }) {
    returning { id }
  }
  insert_steps(objects: $steps) {
    returning { id }
  }
  delete_recipe_images(where: { recipe_id: { _eq: $recipe_id } }) {
    returning { id }
  }
  insert_recipe_images(objects: $images) {
    returning { id }
  }
}
`

func SendUpdateMutation(mutation string, recipe models.Recipe, c *gin.Context) (*models.Response, error) {
	GRAPHQL_API := os.Getenv("GRAPHQL_API")
	if GRAPHQL_API == "" {
		GRAPHQL_API = "http://localhost:8080/v1/graphql"
	}
	client := graphql.NewClient(GRAPHQL_API)

	// Flatten ingredients
	ingredients := make([]map[string]interface{}, len(recipe.Ingredients))
	for i, ri := range recipe.Ingredients {
		ingredients[i] = map[string]interface{}{
			"recipe_id": *recipe.Id,
			"quantity":  ri.Amount,
			"unit":      ri.Unit,
			"ingredient": map[string]interface{}{
				"data": map[string]interface{}{
					"name": ri.Ingredient.Name, // Flattened ingredient name
				},
			},
		}
	}

	// Flatten steps
	steps := make([]map[string]interface{}, len(recipe.Steps))
	for i, rs := range recipe.Steps {
		step := map[string]interface{}{
			"recipe_id":   *recipe.Id,
			"step_number": rs.StepNumber,
			"instruction": rs.Instruction,
			"recipe_images": map[string]interface{}{
				"data": map[string]interface{}{
					"image_url": rs.RecipeImages.ImageURL,
				},
			},
		}
		steps[i] = step
	}
	log.Print("sssssssss", steps)

	// Flatten images
	images := make([]map[string]interface{}, len(recipe.Images))
	for i, ri := range recipe.Images {
		images[i] = map[string]interface{}{
			"recipe_id": *recipe.Id,
			"image_url": ri.ImageURL,
		}
	}

	// Create a request
	req := graphql.NewRequest(mutation)
	req.Header.Set("Authorization", c.GetHeader("Authorization"))

	// Set variables dynamically
	req.Var("recipe_id", *recipe.Id)
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

	// Execute the request
	var response struct {
		UpdateRecipesByPk int `json:"update_recipes_by_pk"`
	}

	err := client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
