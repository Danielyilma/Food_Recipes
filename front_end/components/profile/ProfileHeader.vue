<template>
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
      <div class="flex items-center space-x-6">
        <div class="relative">
          <img 
            :src="config.public.imageDomainPath + user?.profile_image" 
            :alt="user?.username"
            class="w-24 h-24 rounded-full object-cover"
          >
          <button 
            @click="openEditProfileModal"
            class="absolute bottom-0 right-0 bg-primary-600 text-white p-2 rounded-full hover:bg-primary-700"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
            </svg>
          </button>
        </div>
        
        <div class="flex-1">
          <div class="flex items-center justify-between">
            <div>
              <h1 class="text-2xl font-bold text-gray-900 dark:text-white">{{ user.username }}</h1>
              <p class="text-gray-500 dark:text-gray-400">@{{ user.tag_name }}</p>
            </div>
            
            <div class="flex space-x-4">
              <NuxtLink 
                to="/recipes/create"
                class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700"
              >
                Create Recipe
              </NuxtLink>
              
              <button 
                @click="openEditProfileModal"
                class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700"
              >
                Edit Profile
              </button>
            </div>
          </div>
          
          <p class="mt-4 text-gray-600 dark:text-gray-300">{{ user.bio }}</p>
          
          <div class="mt-4 flex items-center space-x-6">
            <div>
              <span class="font-semibold text-gray-900 dark:text-white">{{ user.recipes_aggregate.aggregate.count }}</span>
              <span class="text-gray-500 dark:text-gray-400"> recipes</span>
            </div>
            <!-- <div>
              <span class="font-semibold text-gray-900 dark:text-white">{{ user.followersCount }}</span>
              <span class="text-gray-500 dark:text-gray-400"> followers</span>
            </div>
            <div>
              <span class="font-semibold text-gray-900 dark:text-white">{{ user.followingCount }}</span>
              <span class="text-gray-500 dark:text-gray-400"> following</span>
            </div> -->
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import type { User } from '~/types/user'
  
  defineProps<{
    user: any
  }>()
  
  const emit = defineEmits<{
    (e: 'edit'): void
  }>()

  const config = useRuntimeConfig()
  
  const openEditProfileModal = () => {
    emit('edit')
  }
  </script>