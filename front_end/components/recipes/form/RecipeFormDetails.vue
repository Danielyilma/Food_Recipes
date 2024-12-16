<template>
    <div class="space-y-6">
      <div>
        <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Recipe Details</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Category</label>
            <select
              v-model="category_id"
              @change="$emit('update:category', ($event.target as HTMLSelectElement).value)"
              class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
            >
              <option value="">Select a category</option>
              <option v-for="cat in useCategory.categories" :key="cat.id" :value="String(cat.id)">
                {{ cat.name }}
              </option>
            </select>
          </div>
          
          <!-- <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Difficulty</label>
            <select
              v-model="difficulty"
              @change="$emit('update:difficulty', $event.target.value)"
              class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
            >
              <option value="Easy">Easy</option>
              <option value="Medium">Medium</option>
              <option value="Hard">Hard</option>
            </select>
          </div> -->
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Preparation Time</label>
            <div class="flex gap-2">
              <input
                type="number"
                v-model="prepTimeValue"
                min="1"
                @input="updatePrepTime"
                class="block w-20 rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
              >
              <select
                v-model="prepTimeUnit"
                @change="updatePrepTime"
                class="block rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
              >
                <option value="mins">Minutes</option>
                <option value="hours">Hours</option>
              </select>
            </div>
          </div>
          
          <!-- <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Cooking Time</label>
            <div class="flex gap-2">
              <input
                type="number"
                v-model="cookTimeValue"
                min="1"
                @input="updateCookTime"
                class="block w-20 rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
              >
              <select
                v-model="cookTimeUnit"
                @change="updateCookTime"
                class="block rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
              >
                <option value="mins">Minutes</option>
                <option value="hours">Hours</option>
              </select>
            </div>
          </div> -->
          
          <!-- <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Servings</label>
            <input
              type="number"
              v-model="servings"
              min="1"
              @input="$emit('update:servings', parseInt($event.target.value))"
              class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
            >
          </div> -->
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
          @click="$emit('submit')"
          :disabled="!isValid"
          class="px-4 py-2 bg-primary-600 text-white rounded-lg disabled:opacity-50"
        >
          {{ status }} Recipe
        </button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { categories } from '~/data/categories'
  import type { Recipe } from '~/types/recipe'

  const useCategory = useCategoryStore()
  
  const props = defineProps<{
    category_id: string
    prepTime: string
    status: string
  }>()
  
  const emit = defineEmits<{
    (e: 'update:category', value: string): void
    (e: 'update:prepTime', value: string): void
    (e: 'submit'): void
    (e: 'back'): void
  }>()
  
  const category_id = ref(props.category_id)
  // const difficulty = ref(props.difficulty)
  // const servings = ref(props.servings)
  
  // Time handling
  const prepTimeValue = ref(props.prepTime || "")
  const prepTimeUnit = ref('mins')
  // const cookTimeUnit = ref('mins')
  
  const updatePrepTime = () => {
    emit('update:prepTime', `${prepTimeValue.value}`)
  }
  
  // const updateCookTime = () => {
  //   emit('update:cookTime', `${cookTimeValue.value} ${cookTimeUnit.value}`)
  // }
  
  const isValid = computed(() => 
    category_id.value &&
    prepTimeValue.value
  )
  </script>