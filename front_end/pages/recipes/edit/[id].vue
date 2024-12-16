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
            v-model:category_id="formData.category_id"
            v-model:prep-time="formData.prepTime"
            @submit="handleSubmit"
            @back="previousStep"
            status="Edit"
          />
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted } from 'vue'
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
  
  const userStore = useUserStore()
  const config = useRuntimeConfig()
  const formData = ref({
    images: [] as string[],
    featuredImage: '',
    title: '',
    description: '',
    ingredients: [] as { name: string; amount: string; unit: string }[],
    steps: [] as { description: string; image?: string }[],
    category_id: '1',
    prepTime: '' as string
  })

  const recipeStore = useRecipesStore()
  
  onMounted(async () => {
    try {
      loading.value = true
      await recipeStore.fetchRecipe(recipeId)
      if (recipeStore.recipe) {
        formData.value.category_id = recipeStore.recipe.category_id
        formData.value.featuredImage = recipeStore.recipe.recipe_images[0]?.image_url
        formData.value.title = recipeStore.recipe.title
        formData.value.description = recipeStore.recipe.description
        formData.value.prepTime = recipeStore.recipe.prep_time
        recipeStore.recipe.recipe_ingredients.forEach((ing: any) => {
          const ingredient =  {
            unit: ing.unit,
            amount: ing.quantity,
            name: ing.ingredient.name
          }
          formData.value.ingredients?.push(ingredient)
        })

        recipeStore.recipe.recipe_images.forEach((img: any) => {
          formData.value.images?.push(img?.image_url)
        })

        recipeStore.recipe.steps.forEach((step: any) => {
          formData.value.steps?.push({description: step.instruction, image: step.recipe_images[0]?.image_url})
        })
        console.log(formData, "eeeeeeeeeeeeeeeeeee")
      }
    } catch (err) {
      console.log(err)
      error.value = 'Failed to load recipe. Please try again.'
      console.error('Error fetching recipe:', err)
    } finally {
      loading.value = false
    }
  })
  
  const nextStep = () => {
    console.log("next")
    if (currentStep.value < 4) {
      currentStep.value++
    }
  }
  
  const previousStep = () => {
    console.log("next")
    if (currentStep.value > 1) {
      currentStep.value--
    }
  }
  
  // const handleSubmit = async () => {
  //   try {
  //     loading.value = true
  //     await recipeStore.updateRecipe(recipeId, formData.value)
  //     await navigateTo('/profile')
  //   } catch (err) {
  //     error.value = 'Failed to update recipe. Please try again.'
  //     console.error('Error updating recipe:', err)
  //   } finally {
  //     loading.value = false
  //   }
  // }

  const handleSubmit = async () => {
    try {

      const formsData = new FormData()

      formsData.append("title", formData.value?.title);
      formsData.append("description", formData.value.description);
      formsData.append("category_id", formData.value.category_id);
      formsData.append("prepTime", formData.value.prepTime);
      formsData.append("recipe_id", recipeStore.recipe?.id);

      const image_url = await upload_image(formData.value.featuredImage)
      formsData.append('featured_image', image_url  );

      formData.value.ingredients.forEach((ingredient, index) => {
        formsData.append(`ingredient[${index}].name`, ingredient.name);
        formsData.append(`ingredient[${index}].amount`, ingredient.amount);
        formsData.append(`ingredient[${index}].unit`, ingredient.unit);
      })

      for (const [index, step] of formData.value.steps.entries()) {
          formsData.append(`steps[${index}].description`, step.description);
          formsData.append(`steps[${index}].step_number`, String(index + 1));
          if (step.image) {
            const image_url = await upload_image(step.image)
            formsData.append(`steps[${index}].image`, image_url);
          }
      }

      for (const [index, image] of formData.value.images.entries()) {
        const image_url = await upload_image(image)
        formsData.append(`images[${index}]`, image_url);
      }

      const res = await fetch(config.public.fileUploadApi + "update-recipe", {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${userStore.token}`
        },
        body: formsData
      })

      
      const result = await res.json()
      console.log(result)
      console.log('Submitting recipe:', formData)
      // // Redirect to the recipe page after successful creation
      await navigateTo('/profile')
    } catch (error) {
      console.error('Error creating recipe:', error)
    }
  }

  const upload_image = async (image: string) => {
    const formsData = new FormData()
    formsData.append("file", image);

    const uploadResponse = await fetch(config.public.fileUploadApi + "single-upload", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${userStore.token}`
        },
        body: formsData
    });

      if (!uploadResponse.ok) {
        return "";
      }
      
      const { url } = await uploadResponse.json();
      console.log(url)
      return url || ""
  }
  </script>