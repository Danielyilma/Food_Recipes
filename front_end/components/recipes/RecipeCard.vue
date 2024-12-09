<template>
  <div
    class="bg-white dark:bg-gray-700 rounded-lg overflow-hidden shadow-lg transition-transform hover:scale-105"
  >
    <div class="relative">
      <img
        :src="config.public.imageDomainPath + recipe.thumbnail"
        :alt="recipe.title"
        class="w-full h-48 object-cover"
      />
      <div class="absolute top-4 right-4">
        <div
          class="flex items-center justify-center w-8 h-8 rounded-full bg-white dark:bg-gray-800 shadow-md"
          :title="recipe.isPaid ? 'Premium Recipe' : 'Free Recipe'"
        >
          <svg
            v-if="recipe.isPaid"
            class="w-4 h-4 text-primary-600 dark:text-primary-400"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
          <svg
            v-else
            class="w-4 h-4 text-green-600 dark:text-green-400"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path>
          </svg>
        </div>
      </div>
    </div>

    <!-- <img
      :src="recipe.image"
      :alt="recipe.title"
      class="w-full h-48 object-cover"
    /> -->
    <div class="p-6">
      <div class="flex items-center justify-between mb-2">
        <span
          class="text-sm text-primary-600 dark:text-primary-400 font-medium"
          >{{ recipe.category.name }}</span
        >
        <span class="text-sm text-gray-500 dark:text-gray-400">{{
          getPrepTimeMinutes(recipe.prep_time)
        }}</span>
      </div>
      <h3 class="text-xl font-semibold mb-2 dark:text-white">
        {{ recipe.title }}
      </h3>
      <p class="text-gray-600 dark:text-gray-300 mb-4 line-clamp-2">
        {{ recipe.description }}
      </p>
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <img
            :src="config.public.imageDomainPath + recipe.user.profile_image"
            :alt="recipe.user.username"
            class="w-8 h-8 rounded-full"
          />
          <span class="ml-2 text-sm text-gray-600 dark:text-gray-300">{{
            recipe.user.username
          }}</span>
        </div>
        <NuxtLink
          :to="'/recipes/' + recipe.id"
          class="text-primary-600 hover:text-primary-700 dark:text-primary-400 dark:hover:text-primary-500 font-medium"
        >
          View Details
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Recipe } from "~/types/recipe";
import { getPrepTimeMinutes } from "~/utils/time";

const config = useRuntimeConfig()
defineProps<{
  recipe: Recipe;
}>();
</script>
