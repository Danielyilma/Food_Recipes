<template>
    <div class="relative">
      <input 
        type="search" 
        placeholder="Search recipes..." 
        v-model="searchQuery"
        @input="handleSearch"
        class="w-64 px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-none focus:ring-2 focus:ring-primary-500"
      >
    </div>
  </template>
  
  <script setup lang="ts">
  const searchQuery = ref('')
  
  const handleSearch = () => {
    // Implement search functionality
    console.log('Searching for:', searchQuery.value)
  }
  </script>