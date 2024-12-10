<template>
  <div class="space-y-6">
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
      <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">About</h2>
      <div class="space-y-4">
        <div class="flex items-center text-gray-600 dark:text-gray-300">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M3 5.25C3 4.56 3.56 4 4.25 4h3.18c.67 0 1.26.42 1.47 1.05l.7 2.1a1.25 1.25 0 01-.34 1.3l-1.48 1.48a15.8 15.8 0 006.36 6.36l1.48-1.48c.34-.34.85-.45 1.3-.34l2.1.7c.63.21 1.05.8 1.05 1.47v3.18c0 .69-.56 1.25-1.25 1.25H18C10.82 20 4 13.18 4 6V5.25z" />
          </svg>
          {{ user?.phone_number }}
        </div>

        <div class="flex items-center text-gray-600 dark:text-gray-300">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
              d="M2 6.5C2 5.12 3.12 4 4.5 4h15c1.38 0 2.5 1.12 2.5 2.5v11c0 1.38-1.12 2.5-2.5 2.5h-15C3.12 20 2 18.88 2 17.5v-11z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
              d="M22 6.5L12 13 2 6.5" />
          </svg>
          {{ user?.email }}
        </div>
        
        <div class="flex items-center text-gray-600 dark:text-gray-300">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          Joined {{ formatDate(user?.created_at) }}
        </div>
      </div>
    </div>
    
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
      <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Popular Categories</h2>
      <div class="space-y-2">
        <div v-for="category in categoriesStore.categories" :key="category.name" class="flex items-center justify-between">
          <span class="text-gray-600 dark:text-gray-300">{{ category.name }}</span>
          <span class="text-gray-500 dark:text-gray-400">{{ category?.recipes_aggregate?.aggregate?.count }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { User } from '~/types/user'

const props = defineProps<{
  user: any
}>()
console.log(props.user)
const categoriesStore = useCategoryStore()
categoriesStore.fetchCategories()
// const popularCategories = [
//   { name: 'Breakfast', count: 5 },
//   { name: 'Main Course', count: 8 },
//   { name: 'Desserts', count: 3 }
// ]

const formatDate = (date: string) => {
  console.log(date)
  // date = date?.replace(/\..*/, '')
    return new Date(date).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    })
  }
</script>