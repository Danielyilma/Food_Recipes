<template>
    <div class="flex items-center space-x-4">
      <NuxtLink 
        to="/recipes" 
        class="text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
      >
        Recipes
      </NuxtLink>
      <NuxtLink 
        v-if="userStore.isAuthenticated()"
        to="/bookmarks" 
        class="text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
      >
        Bookmarks
      </NuxtLink>
    </div>
  </template>
  
  <script setup lang="ts">
  const userStore = useUserStore()
  </script>