<template>
    <div class="text-center py-12">
      <div class="inline-block p-4 rounded-full bg-gray-100 dark:bg-gray-700 mb-4">
        <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
        </svg>
      </div>
      
      <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">No bookmarks yet</h2>
      <p class="text-gray-600 dark:text-gray-400 mb-6">Start saving your favorite recipes to access them quickly later.</p>
      
      <NuxtLink
        to="/recipes"
        class="inline-flex items-center px-4 py-2 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700"
      >
        Browse Recipes
        <svg class="ml-2 -mr-1 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </NuxtLink>
    </div>
  </template>