<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
    <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">
      Ingredients
      <span class="text-sm font-normal text-gray-500 dark:text-gray-400">
        ({{ servings }} servings)
      </span>
    </h2>

    <div class="flex items-center space-x-4 mb-6">
      <button
        @click="updateServings(-1)"
        class="p-1 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
        :disabled="servings <= 1"
      >
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            d="M20 12H4"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
          />
        </svg>
      </button>

      <span class="text-lg font-medium text-gray-900 dark:text-white">{{
        servings
      }}</span>

      <button
        @click="updateServings(1)"
        class="p-1 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700"
      >
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            d="M12 4v16m8-8H4"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
          />
        </svg>
      </button>
    </div>

    <ul class="space-y-3">
      <li
        v-for="ingredient in ingredients"
        :key="ingredient.id"
        class="flex items-center text-gray-700 dark:text-gray-300"
      >
        <span class="w-16 font-medium"
          >{{ calculateAmount(ingredient.quantity) }} {{ ingredient.unit }}</span
        >
        <span>{{ ingredient.ingredient.name }}</span>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import type { Ingredient } from "~/types/recipe";

const props = defineProps<{
  ingredients: Ingredient[];
  initialServings: number;
}>();

const servings = ref(props.initialServings);

const updateServings = (change: number) => {
  const newServings = servings.value + change;
  if (newServings >= 1) {
    servings.value = newServings;
  }
};

const calculateAmount = (amount: string) => {
  const baseAmount = parseFloat(amount);
  console.log(baseAmount, amount)
  return ((baseAmount * servings.value) / props.initialServings).toFixed(1);
};
</script>
