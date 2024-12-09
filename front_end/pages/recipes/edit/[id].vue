<template>
    <div class="container mx-auto px-4 py-8">
      <div class="max-w-3xl mx-auto">
        <div class="flex items-center justify-between mb-8">
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white">Edit Recipe</h1>
          <button
            @click="navigateTo('/profile')"
            class="text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
  
        <div v-if="loading" class="flex justify-center py-12">
          <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
        </div>
  
        <div v-else-if="error" class="bg-red-50 dark:bg-red-900/50 p-4 rounded-lg">
          <p class="text-red-600 dark:text-red-400">{{ error }}</p>
        </div>
  
        <div v-else class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
          <RecipeFormProgress :current-step="currentStep" :total-steps="4" class="mb-8" />
  
          <RecipeFormBasics
            v-if="currentStep === 1"
            v-model:title="formData.title"
            v-model:description="formData.description"
            @next="nextStep"
          />
          
          <RecipeFormIngredients
            v-if="currentStep === 2"
            v-model:ingredients="formData.ingredients"
            @next="nextStep"
            @back="previousStep"
          />
          
          <RecipeFormSteps
            v-if="currentStep === 3"
            v-model:steps="formData.steps"
            @next="nextStep"
            @back="previousStep"
          />
          
          <RecipeFormDetails
            v-if="currentStep === 4"
            v-model:category="formData.category"
            v-model:prep-time="formData.prepTime"
            v-model:cook-time="formData.cookTime"
            v-model:servings="formData.servings"
            v-model:difficulty="formData.difficulty"
            @submit="handleSubmit"
            @back="previousStep"
          />
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted } from 'vue'
  import type { Recipe } from '~/types/recipe'
  import RecipeFormBasics from '~/components/recipes/form/RecipeFormBasics.vue'
  import RecipeFormSteps from '~/components/recipes/form/RecipeFormSteps.vue'
  import RecipeFormDetails from '~/components/recipes/form/RecipeFormDetails.vue'
  import RecipeFormIngredients from '~/components/recipes/form/RecipeFormIngredients.vue'
  import RecipeFormProgress from '~/components/recipes/form/RecipeFormProgress.vue'

  definePageMeta({
    middleware: "authentication",
  });
  const route = useRoute()
  const recipeId = parseInt(route.params.id as string)
  const currentStep = ref(1)
  const loading = ref(true)
  const error = ref('')
  
  const formData = ref<Partial<Recipe>>({
    title: '',
    description: '',
    ingredients: [],
    steps: [],
    category: '',
    prepTime: '',
    cookTime: '',
    servings: 4,
    difficulty: 'Easy'
  })
  
  const recipeStore = useRecipesStore()
  
  onMounted(async () => {
    try {
      loading.value = true
      const recipe = await recipeStore.fetchRecipe(recipeId)
      if (recipe) {
        formData.value = { ...recipe }
      }
    } catch (err) {
      error.value = 'Failed to load recipe. Please try again.'
      console.error('Error fetching recipe:', err)
    } finally {
      loading.value = false
    }
  })
  
  const nextStep = () => {
    if (currentStep.value < 4) {
      currentStep.value++
    }
  }
  
  const previousStep = () => {
    if (currentStep.value > 1) {
      currentStep.value--
    }
  }
  
  const handleSubmit = async () => {
    try {
      loading.value = true
      await recipeStore.updateRecipe(recipeId, formData.value)
      await navigateTo('/profile')
    } catch (err) {
      error.value = 'Failed to update recipe. Please try again.'
      console.error('Error updating recipe:', err)
    } finally {
      loading.value = false
    }
  }
  </script>