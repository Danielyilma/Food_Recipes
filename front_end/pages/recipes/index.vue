<template>
  <div class="container mx-auto px-4 py-8">
    <div class="flex flex-col md:flex-row gap-8">
      <!-- Sidebar Filters -->
      <div class="w-full md:w-64 flex-shrink-0">
        <ClientOnly>
          <RecipeFilters />
        </ClientOnly>
      </div>

      <!-- Main Content -->
      <div class="flex-1">
        <div class="mb-8">
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-4">
            Browse Recipes
          </h1>
          <ClientOnly>
            <RecipeSearch/>
          </ClientOnly>
        </div>

        <!-- Results Count -->
        <p class="text-sm text-gray-600 dark:text-gray-400 mb-6">
          Showing {{ useRecipe.recipes.length }} recipes
        </p>

        <!-- Recipe Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <RecipeCard
            v-for="recipe in useRecipe.recipes"
            :key="recipe.id"
            :recipe="recipe"
          />
        </div>

        <!-- Pagination -->
        <!-- <Pagination
          v-if="totalItems > 0"
          :current-page="currentPage"
          :total-items="totalItems"
          :items-per-page="itemsPerPage"
          @update:page="updatePage"
        /> -->

        <!-- No Results -->
        <div v-if="useRecipe.recipes.length === 0" class="text-center py-12">
          <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">
            No recipes found
          </h3>
          <p class="text-gray-600 dark:text-gray-400">
            Try adjusting your search or filters
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import RecipeCard from "~/components/recipes/RecipeCard.vue";
import RecipeFilters from "~/components/recipes/RecipeFilters.vue";
import RecipeSearch from "~/components/recipes/RecipeSearch.vue";
import Pagination from "~/components/ui/Pagination.vue";
import { usePagination } from "~/composables/usePagination";
import { getPrepTimeMinutes } from "~/utils/time";
import type { Filters, Recipe } from "~/types/recipe";


const useRecipe = useRecipesStore();
const useFilter = useRecipeFilters();

await useRecipe.fetchRecipes(1);

// const recipes: Recipe[] = useRecipe.recipes;
// Sample data - In a real app, this would come from an API
// const recipes: Recipe[] = [
//   {
//     id: 1,
//     title: "Classic Pancakes",
//     description: "Fluffy and delicious pancakes perfect for breakfast",
//     image:
//       "https://images.unsplash.com/photo-1528207776546-365bb710ee93?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
//     category: { name: "Breakfast" },
//     prep_time: "20 mins",
//     user: {
//       username: "John Doe",
//       profile_image:
//         "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
//     },
//     isPaid: false,
//   },
//   {
//     id: 2,
//     title: "Grilled Salmon",
//     description: "Healthy and tasty grilled salmon with vegetables",
//     image:
//       "https://images.unsplash.com/photo-1467003909585-2f8a72700288?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
//     category: { name: "Main Course" },
//     prep_time: "35 mins",
//     user: {
//       username: "Jane Smith",
//       profile_image:
//         "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
//     },
//     isPaid: true,
//   },
//   {
//     id: 3,
//     title: "Chocolate Cake",
//     description: "Rich and moist chocolate cake for special occasions",
//     image:
//       "https://images.unsplash.com/photo-1578985545062-69928b1d9587?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
//     category: { name: "Desserts" },
//     prep_time: "75 mins",
//     user: {
//       username: "Mike Johnson",
//       profile_image:
//         "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
//     },
//     isPaid: true,
//   },
// ];

// const { currentPage, paginatedItems, totalItems, itemsPerPage, updatePage } =
//   usePagination(filteredRecipes.value);
</script>

<!-- if (activeFilters.value.prepTime) {
  filtered = filtered.filter((recipe) => {
    const minutes = getPrepTimeMinutes(recipe.prepTime);
    switch (activeFilters.value.prepTime) {
      case "quick":
        return minutes < 30;
      case "medium":
        return minutes >= 30 && minutes <= 60;
      case "long":
        return minutes > 60;
      default:
        return true;
    }
  });
} -->
