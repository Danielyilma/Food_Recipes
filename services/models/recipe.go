package models

// Ingredient represents the input for an ingredient in the recipe.

type Recipe_ingredients struct {
	Ingredient Ingredient `form:"ingredient"`
	Amount     string     `form:"amount"`
	Unit       string     `form:unit`
}

type Ingredient struct {
	Name string `form:"name"`
}

// RecipeImage represents the input for an image associated with a recipe step or the recipe itself.
type RecipeImage struct {
	ImageURL string `form:"image_url"`
}

// Step represents the steps in the recipe, including instructions and associated images.
type Step struct {
	StepNumber   int         `form:"step_number"`
	Instruction  string      `form:"instruction"`
	RecipeImages RecipeImage `form:"recipe_images"`
}

// FileUploadInput represents the input structure for the fileUpload mutation.
type Recipe struct {
	Title       string               `form:"title"`
	Description string               `form:"description"`
	PrepTime    int                  `form:"prepTime"`
	Thumbnail   string               `form:"thumbnail"`
	IsPaid      bool                 `form:"is_paid"`
	Price       float64              `form:"price"`
	UserID      int                  `form:"user_id"`
	CategoryID  int                  `form:"category_id"`
	Ingredients []Recipe_ingredients `form:"ingredients[]"`
	Steps       []Step               `form:"steps[]"`
	Images      []RecipeImage        `form:"images[]"`
}

type InsertRecipesOne struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}
