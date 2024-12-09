<template>
    <div class="space-y-6">
      <div>
        <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Recipe Ingredients</h2>
        
        <div class="space-y-4">
          <div
            v-for="(ingredient, index) in ingredients"
            :key="index"
            class="flex gap-4"
          >
            <div class="flex-1">
              <input
                type="text"
                v-model="ingredient.name"
                @input="updateIngredients"
                class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
                placeholder="Ingredient name"
              >
            </div>
            <div class="w-24">
              <input
                type="text"
                v-model="ingredient.amount"
                @input="updateIngredients"
                class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
                placeholder="Amount"
              >
            </div>
            <div class="w-24">
              <input
                type="text"
                v-model="ingredient.unit"
                @input="updateIngredients"
                class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
                placeholder="Unit"
              >
            </div>
            <button
              @click="removeIngredient(index)"
              class="p-2 text-red-600 hover:text-red-700 dark:text-red-400"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
          
          <button
            @click="addIngredient"
            class="flex items-center text-primary-600 hover:text-primary-700 dark:text-primary-400"
          >
            <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add Ingredient
          </button>
        </div>
      </div>
      
      <div class="flex justify-between">
        <button
          @click="$emit('back')"
          class="px-4 py-2 border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700"
        >
          Back
        </button>
        <button
          @click="$emit('next')"
          :disabled="!isValid"
          class="px-4 py-2 bg-primary-600 text-white rounded-lg disabled:opacity-50"
        >
          Next Step
        </button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  interface Ingredient {
    name: string
    amount: string
    unit: string
  }
  
  const props = defineProps<{
    ingredients: Ingredient[]
  }>()
  
  const emit = defineEmits<{
    (e: 'update:ingredients', value: Ingredient[]): void
    (e: 'next'): void
    (e: 'back'): void
  }>()
  
  const addIngredient = () => {
    emit('update:ingredients', [
      ...props.ingredients,
      { name: '', amount: '', unit: '' }
    ])
  }
  
  const removeIngredient = (index: number) => {
    const newIngredients = [...props.ingredients]
    newIngredients.splice(index, 1)
    emit('update:ingredients', newIngredients)
  }
  
  const updateIngredients = () => {
    emit('update:ingredients', [...props.ingredients])
  }
  
  const isValid = computed(() => 
    props.ingredients.length > 0 &&
    props.ingredients.every(ingredient => 
      ingredient.name.trim() && 
      ingredient.amount.trim()
    )
  )
  </script>