<template>
    <div class="space-y-6">
      <div class="flex items-center justify-between">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white">My Recipes</h2>
        <div class="flex items-center space-x-4">
          <select 
            v-model="sortBy"
            class="rounded-md border-gray-300 dark:border-gray-600 dark:bg-gray-700"
          >
            <option value="newest">Newest First</option>
            <option value="oldest">Oldest First</option>
            <option value="popular">Most Popular</option>
          </select>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div 
          v-for="recipe in sortedRecipes" 
          :key="recipe.id"
          class="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden"
        >
          <img 
            :src="config.public.imageDomainPath + recipe.thumbnail" 
            :alt="recipe.title"
            class="w-full h-48 object-cover"
          >
          
          <div class="p-4">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm text-primary-600 dark:text-primary-400">{{ recipe.category.name }}</span>
              <span class="text-sm text-gray-500 dark:text-gray-400">{{ formatDate(recipe.created_at) }}</span>
            </div>
            
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">{{ recipe.title }}</h3>
            <p class="text-gray-600 dark:text-gray-300 mb-4 line-clamp-2">{{ recipe.description }}</p>
            
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-2">
                <span class="text-gray-600 dark:text-gray-300">{{ recipe.likes_aggregate.aggregate.count }} likes</span>
              </div>
              
              <div class="flex items-center space-x-2">
                <button 
                  @click="editRecipe(recipe.id)"
                  class="p-2 text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                
                <button 
                  @click="deleteRecipe(recipe.id)"
                  class="p-2 text-gray-600 dark:text-gray-300 hover:text-red-600"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { recipes } from '~/data/recipes';
import  { deleteRecipeQuery } from '~/queries/recipe';
import type { Recipe } from '~/types/user'
  
  const props = defineProps<{
    recipes: Recipe[]
  }>()
  
  const emit = defineEmits([
    "update:recipes"
  ])

  const sortBy = ref('newest')
  const config = useRuntimeConfig()
  
  const sortedRecipes = computed(() => {
    return [...props.recipes].sort((a, b) => {
      switch (sortBy.value) {
        case 'oldest':
          return new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
        case 'popular':
          return b.likes_aggregate.aggregate.count - a.likes_aggregate.aggregate.count
        default: // newest
          return new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
      }
    })
  })
  
  const formatDate = (date: string) => {
    return new Date(date).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    })
  }
  
  const editRecipe = (id: number) => {
    navigateTo(`/recipes/edit/${id}`)
  }
  
  const deleteRecipe = async (id: number) => {
    if (confirm('Are you sure you want to delete this recipe?')) {
      await sendMutation(deleteRecipeQuery, {id: id})
      emit('update:recipes', props.recipes.filter(recipe => recipe.id !== id));
      console.log('Deleting recipe:', id)
      return
    }
  }
  </script>