export interface User {
  username: string;
  profile_image: string;
}

export interface Recipe {
  id: number;
  title: string;
  description: string;
  thumbnail: string;
  category: Category;
  prep_time: string;
  user: User;
  isPaid: boolean;
}

export interface RecipeDetail {
  user:               User;
  category:           Category;
  category_id:        number;
  description:        string;
  id:                 number;
  created_at:         Date;
  is_paid:            boolean;
  prep_time:          number;
  price:              number;
  recipe_images:      RecipeImage[];
  steps:              Step[];
  recipe_ingredients: Ingredient[];
}

export interface RecipeImage {
  image_url: string;
}

export interface Filters {
  categories: Category[];
  prepTime: string;
  creator: string;
}

export interface Category {
  id: number;
  name: string;
  description: string;
}

export interface Creator {
  id: number;
  name: string;
  avatar: string;
  recipeCount: number;
}

export interface Ingredient {
  id: number;
  name: string;
  amount: string;
  unit: string;
}

export interface Step {
  id: number;
  step_number: number;
  instruction: string;
  image?: string;
}

export interface Comment {
  id: number;
  user: User;
  content: string;
  createdAt: string;
}

// export interface Recipe {
//   id: number;
//   title: string;
//   description: string;
//   image: string;
//   category: string;
//   prepTime: string;
//   cookTime: string;
//   servings: number;
//   difficulty: "Easy" | "Medium" | "Hard";
//   author: Author;
//   isPaid: boolean;
//   ingredients: Ingredient[];
//   steps: Step[];
//   comments: Comment[];
//   likes: number;
//   isLiked?: boolean;
//   isBookmarked?: boolean;
//   rating: number;
//   userRating?: number;
// }
