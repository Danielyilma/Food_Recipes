<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-4xl mx-auto space-y-8">
      <!-- Recipe Title and Meta -->
      <div class="space-y-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">{{ recipe?.title }}</h1>
            <div class="mt-2 flex items-center space-x-4 text-sm text-gray-500 dark:text-gray-400">
              <span>{{ recipe?.category?.name }}</span>
              <span>•</span>
              <span>{{ recipe?.prep_time }} prep</span>
              <span></span>
              <!-- <span>•</span>
              <span>{{ recipe.cookTime }} cook</span>
              <span>•</span>
              <span>{{ recipe.difficulty }}</span> -->
            </div>
          </div>
          <div class="flex items-center space-x-2">
            <button 
              @click="toggleBookmark"
              class="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
              :class="{ 'text-primary-600 dark:text-primary-400': isBookmarked }"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path 
                  :d="isBookmarked 
                    ? 'M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z'
                    : 'M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z'"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                />
              </svg>
            </button>
            <button 
              @click="toggleLike"
              class="flex p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
              :class="{ 'text-red-600': isLiked }"
            >
              <span class="text-lg dark:text-white pr-2 ml-2">{{ likes }}</span>
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path 
                  :d="isLiked
                    ? 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z'
                    : 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z'"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                />
              </svg>
            </button>
          </div>
        </div>
        
        <!-- Image Gallery -->
        <RecipeImageGallery 
          :images="recipe?.recipe_images" 
          :title="recipe?.title"
        />
        
        <!-- Author Info -->
        <div class="flex items-center space-x-4">
          <img 
            :src="config.public.imageDomainPath + recipe.user?.profile_image" 
            :alt="recipe?.user?.username"
            class="w-12 h-12 rounded-full"
          >
          <div>
            <p class="text-sm font-medium text-gray-900 dark:text-white">
              {{ recipe.user?.username }}
            </p>
            <p class="text-sm text-gray-500 dark:text-gray-400">
              Recipe Creator
            </p>
          </div>
        </div>
        
        <p class="text-gray-600 dark:text-gray-300">
          {{ recipe?.description }}
        </p>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="md:col-span-2">
          <RecipeSteps :steps="recipe?.steps" />
        </div>
        
        <div>
          <RecipeIngredients 
            :ingredients="recipe?.recipe_ingredients"
            :initial-servings="recipe?.servings || 1"
          />
        </div>
      </div>
      
      <RecipeComments :recipe_id="Number(recipeId)" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Recipe, RecipeDetail } from '~/types/recipe'
import RecipeImageGallery from "~/components/recipes/RecipeImageGallery.vue"
import RecipeSteps from "~/components/recipes/RecipeSteps.vue"
import RecipeComments from "~/components/recipes/RecipeComments.vue"
import RecipeIngredients from "~/components/recipes/RecipeIngredients.vue"
import { sendMutation, sendQuery } from '~/utils/req'
import { dislikeRecipeQuery, getlikesByRecipeAndUserQuery, likeRecipeQuery } from '~/queries/recipe'
import { addBookmarksQuery, deleteBookmarkQuery } from '~/queries/bookmarks'

definePageMeta({
  middleware: "authentication",
});

const config = useRuntimeConfig()
const router = useRoute()
const recipeStore = useRecipesStore()
const userStore = useUserStore()
const recipeId = router.params.id

await recipeStore.fetchRecipe(Number(recipeId))
const recipe: RecipeDetail = recipeStore?.recipe

await recipeStore.fetchRecipeLikes(userStore.user.id, Number(recipeId) )

const likes = ref(recipeStore?.recipeLikes)
const variable = {
  recipeId: recipeId,
  userId: userStore?.user?.id
}
const isBookmarked = ref()
const isLiked = ref(recipeStore?.isLiked)
console.log(isLiked.value)
// isBookmarked.value = await sendQuery(addBookmarksQuery, variable)


const toggleBookmark = async () => {
  if (isBookmarked.value) {
    await sendMutation(deleteBookmarkQuery, variable)
  } else {
    await sendMutation(addBookmarksQuery, variable)
  }
  isBookmarked.value = !isBookmarked.value
}

const toggleLike = async () => {
  if (isLiked.value) {
    await sendMutation(dislikeRecipeQuery, variable)
  } else {
    await sendMutation(likeRecipeQuery, variable)
  }
  isLiked.value = !isLiked.value
  likes.value += isLiked.value ? 1 : -1
  // In a real app, this would make an API call
}
</script>