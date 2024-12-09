<template>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="recipe in recipes"
        :key="recipe.recipe.id"
        class="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden group"
      >
        <div class="relative">
          <img 
            :src="recipe?.recipe?.thumbnail" 
            :alt="recipe.recipe.title"
            class="w-full h-48 object-cover"
          >
          <button
            @click="$emit('remove', recipe.recipe.id)"
            class="absolute top-2 right-2 p-2 bg-white dark:bg-gray-800 rounded-full shadow-md opacity-0 group-hover:opacity-100 transition-opacity"
          >
            <svg class="w-5 h-5 text-gray-600 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
          </button>
        </div>
        
        <div class="p-4">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm text-primary-600 dark:text-primary-400">{{ recipe?.recipe.category }}</span>
            <span class="text-sm text-gray-500 dark:text-gray-400">{{ recipe?.recipe.prep_time }}</span>
          </div>
          
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">{{ recipe?.recipe.title }}</h3>
          <p class="text-gray-600 dark:text-gray-300 mb-4 line-clamp-2">{{ recipe?.recipe.description }}</p>
          
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <img 
                :src="recipe.user.profile_image" 
                :alt="recipe.user.username"
                class="w-8 h-8 rounded-full"
              >
              <span class="ml-2 text-sm text-gray-600 dark:text-gray-300">{{ recipe.user.username }}</span>
            </div>
            
            <NuxtLink 
              :to="`/recipes/${recipe.id}`"
              class="text-primary-600 hover:text-primary-700 dark:text-primary-400 dark:hover:text-primary-500 font-medium"
            >
              View Recipe
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import type { Recipe } from '~/types/recipe'
  
  defineProps<{
    recipes: Recipe[]
  }>()
  
  defineEmits<{
    (e: 'remove', id: number): void
  }>()
  </script>