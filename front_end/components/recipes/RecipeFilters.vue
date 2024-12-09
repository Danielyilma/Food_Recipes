<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 space-y-6">
    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Filters</h3>

    <!-- Categories -->
    <div>
      <h4 class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Categories
      </h4>
      <div class="space-y-2">
        <label
          v-for="category in useCategories.categories"
          :key="category.id"
          class="flex items-center"
        >
          <input
            type="checkbox"
            :value="category.name"
            v-model="useFilter.activeFilters.value.categories"
            class="rounded border-gray-300 text-primary-600 focus:ring-primary-500"
          />
          <span class="ml-2 text-sm text-gray-600 dark:text-gray-400">{{
            category.name
          }}</span>
        </label>
      </div>
    </div>

    <!-- Preparation Time -->
    <div>
      <h4 class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Preparation Time
      </h4>
      <select
        v-model="useFilter.activeFilters.value.prepTime"
        class="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-600 dark:bg-gray-700 shadow-sm focus:border-primary-500 focus:ring-primary-500"
      >
        <option value="">Any time</option>
        <option value="quick">Quick (< 30 mins)</option>
        <option value="medium">Medium (30-60 mins)</option>
        <option value="long">Long (> 60 mins)</option>
      </select>
    </div>

    <!-- Creator -->
    <!-- <div>
      <h4 class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Creator
      </h4>
      <select
        v-model="selectedCreator"
        class="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-600 dark:bg-gray-700 shadow-sm focus:border-primary-500 focus:ring-primary-500"
        @change="emitFilters"
      >
        <option value="">All creators</option>
        <option
          v-for="creator in creators"
          :key="creator.id"
          :value="creator.id"
        >
          {{ creator.name }}
        </option>
      </select>
    </div> -->

    <!-- Creator -->
    <div>
      <h4 class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Creator
      </h4>
      <div class="creator-search">
        <CreatorSearch
          v-model="useFilter.activeFilters.value.creator"
          @update:modelValue="handleCreatorChange"
        />
      </div>
    </div>

    <!-- Clear Filters -->
    <button
      @click="useFilter.clearFilters"
      class="w-full px-4 py-2 text-sm font-medium text-primary-600 dark:text-primary-400 hover:bg-primary-50 dark:hover:bg-gray-700 rounded-md"
    >
      Clear Filters
    </button>
  </div>
</template>

<script setup lang="ts">
import CreatorSearch from "./CreatorSearch.vue";
import type { Category, User } from "~/types/recipe"; 

const useFilter = useRecipeFilters();
const useCategories = useCategoryStore()

useCategories.fetchCategories()

const selectedCategories = ref<string[]>([]);
const selectedPrepTime = ref("");
const selectedCreator = ref("");



const handleCreatorChange = (creator: string) => {
  selectedCreator.value = creator;
};
</script>




// const categories: Category[] = [
//   { id: 1, name: "Breakfast" },
//   { id: 2, name: "Main Course" },
//   { id: 3, name: "Desserts" },
//   { id: 4, name: "Healthy" },
//   { id: 5, name: "Vegetarian" },
// ];

// const creators: User[] = [
//   { id: 1, name: "John Doe" },
//   { id: 2, name: "Jane Smith" },
//   { id: 3, name: "Mike Johnson" },
// ];