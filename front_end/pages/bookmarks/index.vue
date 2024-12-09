<template>
    <div class="container mx-auto px-4 py-8">
      <div class="max-w-6xl mx-auto">
        <BookmarksHeader :total-bookmarks="useBookmark.bookmarks.length" />
        
        <div v-if="loading" class="flex justify-center py-12">
          <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
        </div>
        
        <div v-else-if="error" class="bg-red-50 dark:bg-red-900/50 p-4 rounded-lg">
          <p class="text-red-600 dark:text-red-400">{{ error }}</p>
        </div>
        
        <div v-else>
          <BookmarksList 
            v-if="useBookmark.bookmarks.length > 0"
            :recipes="useBookmark.bookmarks"
            @remove="removeBookmark"
          />
          
          <BookmarksEmpty v-else />
          
          <!-- <Pagination
            v-if="bookmarkedRecipes.length > itemsPerPage"
            :current-page="currentPage"
            :total-items="bookmarkedRecipes.length"
            :items-per-page="itemsPerPage"
            @update:page="updatePage"
          /> -->
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import type { Recipe } from '~/types/recipe'
  import { usePagination } from '~/composables/usePagination'
  
  const loading = ref(false)
  const error = ref('')
  const itemsPerPage = 12
  const useBookmark = useBookmarksStore()
  const userStore = useUserStore()

  await useBookmark.fetchBookmarks(userStore.user.id)


  
  // In a real app, this would be fetched from an API
  // const bookmarkedRecipes = ref<Recipe[]>([
  //   {
  //     id: 1,
  //     title: 'Classic Pancakes',
  //     description: 'Fluffy and delicious pancakes perfect for breakfast',
  //     image: 'https://images.unsplash.com/photo-1528207776546-365bb710ee93',
  //     category: 'Breakfast',
  //     prepTime: '20 mins',
  //     cookTime: '15 mins',
  //     servings: 4,
  //     difficulty: 'Easy',
  //     author: {
  //       name: 'John Doe',
  //       avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e'
  //     },
  //     isPaid: false,
  //     ingredients: [],
  //     steps: [],
  //     comments: [],
  //     likes: 42,
  //     isLiked: false,
  //     isBookmarked: true,
  //     rating: 4.5
  //   }
  // ])
  
  // const {
  //   currentPage,
  //   paginatedItems: paginatedRecipes,
  //   updatePage
  // } = usePagination(bookmarkedRecipes.value, itemsPerPage)
  
  const removeBookmark = async (recipeId: number) => {
    try {
      // In a real app, this would make an API call
      useBookmark.bookmarks = useBookmark.bookmarks.filter(recipe => recipe.id !== recipeId)
    } catch (err) {
      error.value = 'Failed to remove bookmark. Please try again.'
    }
  }
  </script>