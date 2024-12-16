<template>
    <div class="p-6">
      <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-6">Purchase Recipe</h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <!-- Recipe Image -->
        <div class="relative aspect-video rounded-lg overflow-hidden">
          <img 
            :src="config.public.imageDomainPath + recipe.thumbnail" 
            :alt="recipe.title"
            class="absolute inset-0 w-full h-full object-cover"
          >
        </div>
        
        <!-- Recipe Details -->
        <div class="space-y-4">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">{{ recipe.title }}</h3>
          
          <div class="flex items-center space-x-4 text-sm text-gray-500 dark:text-gray-400">
            <span>{{ recipe.category.name }}</span>
            <span>•</span>
            <span>{{ recipe.prep_time }} prep</span>
            <!-- <span>•</span>
            <span>{{ recipe.difficulty }}</span> -->
          </div>
          
          <p class="text-gray-600 dark:text-gray-300">{{ recipe.description }}</p>
          
          <div class="flex items-center space-x-4">
            <img 
              :src="config.public.imageDomainPath + recipe.user.profile_image" 
              :alt="recipe.user.username"
              class="w-10 h-10 rounded-full"
            >
            <div>
              <p class="text-sm font-medium text-gray-900 dark:text-white">
                {{ recipe?.user.username }}
              </p>
              <p class="text-sm text-gray-500 dark:text-gray-400">Recipe Creator</p>
            </div>
          </div>
          
          <div class="pt-4 border-t border-gray-200 dark:border-gray-700">
            <div class="flex items-center justify-between mb-4">
              <span class="text-lg font-medium text-gray-900 dark:text-white">Price</span>
              <span class="text-2xl font-bold text-primary-600 dark:text-primary-400">
                ${{ recipe?.price.toFixed(2) }}
              </span>
            </div>
            
            <button
              @click="$emit('continue')"
              class="w-full py-3 px-4 bg-primary-600 hover:bg-primary-700 text-white rounded-lg font-medium"
            >
              Continue to Payment
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import type { Recipe } from '~/types/recipe'
  
  const config = useRuntimeConfig()
  defineProps<{
    recipe: Recipe
  }>()
  
  defineEmits<{
    (e: 'continue'): void
  }>()
  </script>