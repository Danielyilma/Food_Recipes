<template>
    <div class="flex items-center space-x-4">
      <template v-if="userStore.isAuthenticated()">
        <UserMenu />
      </template>
      <template v-else>
        <NuxtLink 
          to="/login"
          class="text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
        >
          Login
        </NuxtLink>
      </template>
    </div>
  </template>
  
  <script setup lang="ts">
  const userStore = useUserStore()
  </script>