<template>
  <div
    class="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-gray-50 dark:bg-gray-900"
  >
    <div
      class="max-w-md w-full space-y-8 bg-white dark:bg-gray-800 p-8 rounded-xl shadow-lg"
    >
      <div class="text-center">
        <h2 class="text-3xl font-bold text-gray-900 dark:text-white">
          {{ title }}
        </h2>
        <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
          <slot name="subtitle"></slot>
        </p>
      </div>

      <slot></slot>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
});
</script>
