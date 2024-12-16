<template>
    <div class="relative">
      <div class="flex items-center justify-between">
        <div 
          v-for="(step, index) in steps" 
          :key="step.id"
          class="flex items-center"
        >
          <div 
            class="flex items-center justify-center w-10 h-10 rounded-full"
            :class="[
              currentStep > index 
                ? 'bg-primary-600 text-white' 
                : currentStep === index + 1
                  ? 'bg-primary-100 dark:bg-primary-800 text-primary-600 dark:text-primary-400'
                  : 'bg-gray-100 dark:bg-gray-700 text-gray-400 dark:text-gray-500'
            ]"
          >
            <span class="text-sm font-medium">{{ index + 1 }}</span>
          </div>
          <span 
            class="ml-3 text-sm font-medium"
            :class="[
              currentStep > index 
                ? 'text-primary-600 dark:text-primary-400' 
                : currentStep === index + 1
                  ? 'text-gray-900 dark:text-white'
                  : 'text-gray-400 dark:text-gray-500'
            ]"
          >
            {{ step.name }}
          </span>
          
          <!-- Connector Line -->
          <div 
            v-if="index < steps.length - 1"
            class="flex-1 mx-4 h-0.5"
            :class="[
              currentStep > index + 1 
                ? 'bg-primary-600' 
                : 'bg-gray-200 dark:bg-gray-700'
            ]"
          ></div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  defineProps<{
    currentStep: number
  }>()
  
  const steps = [
    { id: 1, name: 'Review' },
    { id: 2, name: 'Payment' },
    { id: 3, name: 'Confirmation' }
  ]
  </script>