<template>
    <div class="relative">
      <div class="overflow-hidden h-2 mb-4 text-xs flex rounded bg-gray-200 dark:bg-gray-700">
        <div
          class="shadow-none flex flex-col text-center whitespace-nowrap text-white justify-center bg-primary-600"
          :style="{ width: `${(currentStep / totalSteps) * 100}%` }"
        ></div>
      </div>
      <div class="flex justify-between text-sm">
        <span class="text-gray-600 dark:text-gray-400">Step {{ currentStep }} of {{ totalSteps }}</span>
        <span class="text-gray-600 dark:text-gray-400">{{ stepLabels[currentStep - 1] }}</span>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  defineProps<{
    currentStep: number
    totalSteps: number
  }>()
  
  const stepLabels = [
    'Images',
    'Basic Information',
    'Ingredients',
    'Instructions',
    'Details'
  ]
  </script>