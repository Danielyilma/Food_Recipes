<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
          {{ recipe.title }}
        </h1>
        <div
          class="mt-2 flex items-center space-x-4 text-sm text-gray-500 dark:text-gray-400"
        >
          <span>{{ recipe.category }}</span>
          <span>•</span>
          <span>{{ recipe.prepTime }} prep</span>
          <span>•</span>
          <span>{{ recipe.cookTime }} cook</span>
          <span>•</span>
          <span>{{ recipe.difficulty }}</span>
        </div>
      </div>
      <div class="flex items-center space-x-2">
        <button
          @click="$emit('toggleBookmark')"
          class="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
          :class="{
            'text-primary-600 dark:text-primary-400': recipe.isBookmarked,
          }"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              :d="
                recipe.isBookmarked
                  ? 'M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z'
                  : 'M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z'
              "
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
            />
          </svg>
        </button>
        <button
          @click="$emit('toggleLike')"
          class="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
          :class="{ 'text-red-600': recipe.isLiked }"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              :d="
                recipe.isLiked
                  ? 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z'
                  : 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z'
              "
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
            />
          </svg>
        </button>
      </div>
    </div>

    <div class="relative h-96 rounded-lg overflow-hidden">
      <img
        :src="recipe.image"
        :alt="recipe.title"
        class="absolute inset-0 w-full h-full object-cover"
      />
    </div>

    <div class="flex items-center space-x-4">
      <img
        :src="recipe.author.avatar"
        :alt="recipe.author.name"
        class="w-12 h-12 rounded-full"
      />
      <div>
        <p class="text-sm font-medium text-gray-900 dark:text-white">
          {{ recipe.author.name }}
        </p>
        <p class="text-sm text-gray-500 dark:text-gray-400">Recipe Creator</p>
      </div>
    </div>

    <p class="text-gray-600 dark:text-gray-300">
      {{ recipe.description }}
    </p>
  </div>
</template>

<script setup lang="ts">
import type { Recipe } from "~/types/recipe";

defineProps<{
  recipe: Recipe;
}>();

defineEmits<{
  (e: "toggleBookmark"): void;
  (e: "toggleLike"): void;
}>();
</script>
