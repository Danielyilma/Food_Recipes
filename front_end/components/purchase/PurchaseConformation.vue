<template>
    <div class="p-6 text-center">
      <div class="mb-6">
        <div class="mx-auto w-16 h-16 rounded-full bg-green-100 dark:bg-green-900 flex items-center justify-center">
          <svg class="w-8 h-8 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>
      </div>
      
      <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-4">
        Thank You for Your Purchase!
      </h2>
      
      <p class="text-gray-600 dark:text-gray-300 mb-6">
        You now have access to "{{ recipe.title }}"
      </p>
      
      <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4 mb-6">
        <p class="text-sm text-gray-600 dark:text-gray-300 mb-2">
          Transaction ID: {{ transactionId }}
        </p>
        <p class="text-sm text-gray-600 dark:text-gray-300">
          A receipt has been sent to your email
        </p>
      </div>
      
      <button
        @click="$emit('finish')"
        class="px-6 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700"
      >
        View Recipe
      </button>
    </div>
  </template>
  
  <script setup lang="ts">
  import type { Recipe } from '~/types/recipe'
  
  defineProps<{
    recipe: Recipe
    transactionId: string
  }>()
  
  defineEmits<{
    (e: 'finish'): void
  }>()
  </script>