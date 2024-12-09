<template>
  <div class="space-y-8">
    <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
      Instructions
    </h2>

    <div class="space-y-6">
      <div v-for="(step, index) in steps" :key="step.id" class="flex space-x-6">
        <div class="flex-shrink-0">
          <div
            class="w-8 h-8 rounded-full bg-primary-100 dark:bg-primary-800 flex items-center justify-center"
          >
            <span class="text-primary-600 dark:text-primary-400 font-medium">{{
              index + 1
            }}</span>
          </div>
        </div>

        <div class="flex-1 space-y-4">
          <p class="text-gray-700 dark:text-gray-300">{{ step.instruction }}</p>

          <img
            v-if="step.image"
            :src="config.public.imageDomainPath + step.image[0].image_url"
            :alt="`Step ${step.step_number}`"
            class="rounded-lg w-full h-48 object-cover"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Step } from "~/types/recipe";

const config = useRuntimeConfig()
defineProps<{
  steps: Step[];
}>();
</script>
