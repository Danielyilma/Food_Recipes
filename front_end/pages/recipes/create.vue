<template>
    <ClientOnly>
      <div class="container mx-auto px-4 py-8">
        <div class="max-w-3xl mx-auto">
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-8">Create Recipe</h1>
          
          <!-- Progress Bar -->
          <div class="mb-8">
            <RecipeFormProgress :current-step="currentStep" :total-steps="5" />
          </div>
          
          <!-- Step Content -->
          <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
            <RecipeFormImages
              v-if="currentStep === 1"
              v-model:images="formData.images"
              v-model:featured-image="formData.featuredImage"
              @next="nextStep"
            />
            
            <RecipeFormBasics
              v-if="currentStep === 2"
              v-model:title="formData.title"
              v-model:description="formData.description"
              @next="nextStep"
              @back="previousStep"
            />
            
            <RecipeFormIngredients
              v-if="currentStep === 3"
              v-model:ingredients="formData.ingredients"
              @next="nextStep"
              @back="previousStep"
            />
            
            <RecipeFormSteps
              v-if="currentStep === 4"
              v-model:steps="formData.steps"
              @next="nextStep"
              @back="previousStep"
            />
            
            <RecipeFormDetails
              v-if="currentStep === 5"
              v-model:category="formData.category_id"
              v-model:prep-time="formData.prepTime"
              @submit="handleSubmit"
              @back="previousStep"
            />
          </div>
        </div>
      </div>
    </ClientOnly>
  </template>
  
  <script setup lang="ts">
  import { ref, reactive } from 'vue'
  import type { Recipe } from '~/types/recipe'
  import RecipeFormDetails from "~/components/recipes/form/RecipeFormDetails.vue"
  import RecipeFormSteps from "~/components/recipes/form/RecipeFormSteps.vue"
  import RecipeFormIngredients from "~/components/recipes/form/RecipeFormIngredients.vue"
  import RecipeFormBasics from "~/components/recipes/form/RecipeFormBasics.vue"
  import RecipeFormImages from "~/components/recipes/form/RecipeFormImages.vue"
  import RecipeFormProgress from "~/components/recipes/form/RecipeFormProgress.vue"
    
  const useCategories = useCategoryStore()
  const userStore = useUserStore()
  const config = useRuntimeConfig()

  definePageMeta({
    middleware: "authentication",
  });
  useCategories.fetchCategories()
  const currentStep = ref(1)
  const formData = reactive({
    images: [] as string[],
    featuredImage: '',
    title: '',
    description: '',
    ingredients: [] as { name: string; amount: string; unit: string }[],
    steps: [] as { description: string; image?: string }[],
    category_id: '1',
    prepTime: '' as string
  })
  
  const nextStep = () => {
    if (currentStep.value < 5) {
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

      const formsData = new FormData()

      formsData.append("title", formData.title);
      formsData.append("description", formData.description);
      // formsData.append("prep_time", formData.prepTime);
      formsData.append("category_id", formData.category_id);
      // formsData.append("difficulty", formData.difficulty);
      formsData.append("prepTime", formData.prepTime);
      // formsData.append("servings", formData.servings);

      const image_url = await upload_image(formData.featuredImage)
      formsData.append('featured_image', image_url  );

      formData.ingredients.forEach((ingredient, index) => {
        formsData.append(`ingredient[${index}].name`, ingredient.name);
        formsData.append(`ingredient[${index}].amount`, ingredient.amount);
        formsData.append(`ingredient[${index}].unit`, ingredient.unit);
      })

      for (const [index, step] of formData.steps.entries()) {
          formsData.append(`steps[${index}].description`, step.description);
          formsData.append(`steps[${index}].step_number`, String(index + 1));
          if (step.image) {
            const image_url = await upload_image(step.image)
            formsData.append(`steps[${index}].image`, image_url);
          }
      }

      for (const [index, image] of formData.images.entries()) {
        const image_url = await upload_image(image)
        formsData.append(`images[${index}]`, image_url);
      }

      const res = await fetch(config.public.fileUploadApi + "upload", {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${userStore.token}`
        },
        body: formsData
      })

      
      // const result = await res.json()
      // console.log(result)
      // console.log('Submitting recipe:', formData)
      // // Redirect to the recipe page after successful creation
      // await navigateTo('/recipes')
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
  
  definePageMeta({
    layout: 'default'
  })
  </script>



<!-- v-model:cook-time="formData.cookTime"
v-model:servings="formData.servings"
v-model:difficulty="formData.difficulty" -->