<template>
    <div class="container mx-auto px-4 py-8">
      <div class="max-w-4xl mx-auto">
        <!-- Purchase Progress -->
        <PurchaseProgress :current-step="currentStep" />
        
        <!-- Main Content -->
        <div class="mt-8 bg-white dark:bg-gray-800 rounded-lg shadow-lg overflow-hidden">
          <!-- Recipe Preview -->
          <div v-if="currentStep === 1">
            <PurchaseRecipePreview 
              :recipe="recipeStore.recipe"
              @continue="nextStep"
            />
          </div>
          
          <!-- Payment Form -->
          <div v-if="currentStep === 2">
            <PurchasePaymentForm
              :amount="recipeStore.recipe.price"
              @back="previousStep"
              @submit="handlePayment"
            />
          </div>
          
          <!-- Confirmation -->
          <!-- <div v-if="currentStep === 3">
            <PurchaseConfirmation
              :recipe="recipeStore.recipe"
              :transaction-id="transactionId"
              @finish="navigateToRecipe"
            />
          </div> -->
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import type { Recipe } from '~/types/recipe'
  
  const route = useRoute()
  const recipeStore = useRecipesStore()
  const config = useRuntimeConfig()
  const recipeId = parseInt(route.params.id as string)
  const currentStep = ref(1)
  const transactionId = ref('')
  
  await recipeStore.fetchRecipe(Number(recipeId))
  // In a real app, fetch this from an API
  const recipe = ref<Recipe>({
    id: recipeId,
    title: 'Gourmet Chocolate Cake',
    description: 'A rich and decadent chocolate cake perfect for special occasions',
    price: 9.99,
    image: 'https://images.unsplash.com/photo-1578985545062-69928b1d9587',
    category: 'Desserts',
    prepTime: '45 mins',
    cookTime: '35 mins',
    difficulty: 'Medium',
    author: {
      name: 'Chef John',
      avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e'
    },
    isPaid: true,
    ingredients: [],
    steps: [],
    comments: [],
    likes: 128,
    rating: 4.8,
    servings: 8
  })
  
  const nextStep = () => {
    if (currentStep.value < 3) {
      currentStep.value++
    }
  }
  
  const previousStep = () => {
    if (currentStep.value > 1) {
      currentStep.value--
    }
  }
  
  const handlePayment = async (paymentDetails: any) => {
    try {
      const payload = {
        return_url: `http://localhost:3000/recipes/${recipeStore.recipe.id}`,
        amount: "10",
        phone_number: recipeStore.recipe.user.phone_number
      }

      const response = await fetch(config.public.fileUploadApi + "payment", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
          Authorization: `Bearer ${useUserStore().token}`
        },
        body: JSON.stringify(payload)
      })
      const res = await response.json()
      console.log(res)
      const checkout_url = JSON.parse(res.success)?.data?.checkout_url
      navigateTo(checkout_url, {external: true})
      console.log('Processing payment:', paymentDetails)
      transactionId.value = 'TXN' + Date.now()
      nextStep()
    } catch (error) {
      console.error('Payment failed:', error)
    }
  }
  
  const navigateToRecipe = () => {
    navigateTo(`/recipes/${recipeId}`)
  }
  </script>