<template>
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">My Bookmarks</h1>
        <p class="mt-2 text-gray-600 dark:text-gray-400">
          {{ totalBookmarks }} saved {{ totalBookmarks === 1 ? 'recipe' : 'recipes' }}
        </p>
      </div>
      
      <div class="flex items-center space-x-4">
        <RecipeSearch v-model="searchQuery" />
        <select
          v-model="sortBy"
          class="rounded-md border-gray-300 dark:border-gray-600 dark:bg-gray-700 text-gray-900 dark:text-white"
        >
          <option value="newest">Newest First</option>
          <option value="oldest">Oldest First</option>
        </select>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  
  const props = defineProps<{
    totalBookmarks: number
    bookmarks: any
  }>()
  
  const searchQuery = ref('')
  const sortBy = ref('newest')
  const BookmarkStore = useBookmarksStore()

  const sortedBookmarks = computed(() => {
    return [...props.bookmarks].sort((a, b) => {
      switch (sortBy.value) {
        case 'oldest':
          return new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
        default: // newest
          return new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
      }
    })
  })

  BookmarkStore.bookmarks = sortedBookmarks.value
  </script>