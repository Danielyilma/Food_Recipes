package query

import (
	"context"

	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

var getRecipeQuery = `query GetRecipe($id: Int!) {
		recipes_by_pk(id: $id) {
			id
			title
			description
			prep_time
			thumbnail
			is_paid
			price
			category_id
			recipe_ingredients {
				quantity
				unit
				ingredient {
					name
				}
			}
			steps {
				step_number
        instruction
        recipe_images {
          image_url
        }
			}
			recipe_images {
        image_url
      }
		}
	}`

func GetRecipe(recipeID string, c *gin.Context) (*models.Recipe, error) {
	client := graphql.NewClient("http://localhost:8080/v1/graphql")

	// Create a request
	req := graphql.NewRequest(getRecipeQuery)
	req.Header.Set("Authorization", c.GetHeader("Authorization"))
	req.Var("id", recipeID)

	// Define the response structure
	var response struct {
		RecipesByPK *models.Recipe `json:"recipes_by_pk"`
	}

	// Execute the request
	err := client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	}

	return response.RecipesByPK, nil
}
