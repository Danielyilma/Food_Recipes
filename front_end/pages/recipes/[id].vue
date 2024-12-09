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
              <span class="text-sm ml-2">{{ likes }}</span>
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
              class="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
              :class="{ 'text-red-600': isLiked }"
            >
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
            :initial-servings="recipe?.servings"
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
import { addBookmarksQuery } from '~/queries/bookmarks'

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

await recipeStore.fetchRecipeLikes(Number(recipeId), userStore.user.id)

const likes = ref(recipeStore?.recipeLikes)
const variable = {
  recipeId: recipeId,
  userId: userStore?.user?.id
}
const isBookmarked = ref()
const isLiked = ref(recipeStore?.isLiked)
isLiked.value = await sendQuery(getlikesByRecipeAndUserQuery, variable)
console.log(isLiked.value)
// isBookmarked.value = await sendQuery(addBookmarksQuery, variable)





// In a real app, this would fetch from an API based on the route param
// const recipe: Recipe = {
//   id: 1,
//   title: 'Classic Pancakes',
//   description: 'Light and fluffy pancakes that are perfect for breakfast or brunch. Serve with maple syrup and fresh berries.',
//   recipe_images: [
//     {image_url: 'https://images.unsplash.com/photo-1528207776546-365bb710ee93?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80'}
//   ],
//   category: {
//     name: "Breakfast"
//   },
//   prep_time: '15 mins',
//   cookTime: '20 mins',
//   servings: 4,
//   difficulty: 'Easy',
//   user: {
//     username: 'John Doe',
//     profile_image: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80'
//   },
//   isPaid: false,
//   recipe_ingredients: [
//     { id: 1, ingredient: {name: 'All-purpose flour'}, amount: '1.5', unit: 'cups' },
//     { id: 2, ingredient: {name: 'Baking powder'}, amount: '3.5', unit: 'tsp' },
//     { id: 3, ingredient: {name: 'Salt'}, amount: '0.25', unit: 'tsp' },
//     { id: 4, ingredient: {name: 'Sugar'}, amount: '1', unit: 'tbsp' },
//     { id: 5, ingredient: {name: 'Milk'}, amount: '1.25', unit: 'cups' },
//     { id: 6, ingredient: {name: 'Egg'}, amount: '1', unit: '' },
//     { id: 7, ingredient: {name: 'Melted butter'}, amount: '3', unit: 'tbsp' }
//   ],
//   steps: [
//     {
//       id: 1,
//       step_number: 1,
//       instruction: 'In a large bowl, whisk together the flour, baking powder, salt, and sugar.',
//       image: [
//         {image_url: 'https://images.unsplash.com/photo-1558961363-fa8fdf82db35?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80'}
//       ]
//     },
//     {
//       id: 2,
//       step_number: 2,
//       instruction: 'In another bowl, whisk together the milk, egg, and melted butter.',
//       image: [
//         {image_url: 'https://images.unsplash.com/photo-1568254183919-78a4f43a2877?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80'}
//       ]
//     },
//     {
//       id: 3,
//       step_number: 3,
//       instruction: 'Pour the wet ingredients into the dry ingredients and whisk until just combined. Do not overmix.',
//       image: [
//         {image_url: 'https://images.unsplash.com/photo-1519420573924-65fcd45245f8?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80'}
//       ]
//     }
//   ],
//   comments: [
//     {
//       id: 1,
//       author: {
//         name: 'Jane Smith',
//         avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80'
//       },
//       content: 'These pancakes turned out amazing! So fluffy and delicious.',
//       createdAt: '2024-03-15T08:00:00.000Z'
//     }
//   ],
//   likes: 42,
//   isLiked: false,
//   isBookmarked: false,
//   rating: 4.5
// }

// Combine all recipe images including featured and step images
// const recipeImages = computed(() => {
//   const images = [recipe.image]
//   recipe.steps.forEach(step => {
//     if (step.image) {
//       images.push(step.image)
//     }
//   })
//   return images
// })

const toggleBookmark = async () => {
  await sendMutation(addBookmarksQuery, variable)
  isBookmarked.value = !isBookmarked.value
}

const toggleLike = async () => {
  if (isLiked.value) {
    console.log("disliked")
    await sendMutation(dislikeRecipeQuery, variable)
  } else {
    console.log("liked")
    await sendMutation(likeRecipeQuery, variable)
  }
  isLiked.value = !isLiked.value
  likes.value += isLiked.value ? 1 : -1
  // In a real app, this would make an API call
}
</script>