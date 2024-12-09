<template>
  <nav
    class="flex items-center justify-between border-t border-gray-200 dark:border-gray-700 px-4 py-3 sm:px-6"
    aria-label="Pagination"
  >
    <div class="hidden sm:block">
      <p class="text-sm text-gray-700 dark:text-gray-300">
        Showing
        <span class="font-medium">{{ startIndex + 1 }}</span>
        to
        <span class="font-medium">{{ Math.min(endIndex, totalItems) }}</span>
        of
        <span class="font-medium">{{ totalItems }}</span>
        results
      </p>
    </div>
    <div class="flex-1 flex justify-center sm:justify-end">
      <div class="flex space-x-2">
        <button
          @click="$emit('update:page', currentPage - 1)"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 dark:border-gray-600 text-sm font-medium rounded-md text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Previous
        </button>
        <button
          v-for="pageNumber in visiblePages"
          :key="pageNumber"
          @click="$emit('update:page', pageNumber)"
          :class="[
            'relative inline-flex items-center px-4 py-2 border text-sm font-medium rounded-md',
            currentPage === pageNumber
              ? 'z-10 bg-primary-600 border-primary-600 text-white'
              : 'border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600',
          ]"
        >
          {{ pageNumber }}
        </button>
        <button
          @click="$emit('update:page', currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 dark:border-gray-600 text-sm font-medium rounded-md text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Next
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
const props = defineProps<{
  currentPage: number;
  totalItems: number;
  itemsPerPage: number;
}>();

defineEmits<{
  (e: "update:page", page: number): void;
}>();

const totalPages = computed(() =>
  Math.ceil(props.totalItems / props.itemsPerPage)
);
const startIndex = computed(() => (props.currentPage - 1) * props.itemsPerPage);
const endIndex = computed(() => startIndex.value + props.itemsPerPage);

const visiblePages = computed(() => {
  const pages: number[] = [];
  const maxVisiblePages = 5;
  let start = Math.max(1, props.currentPage - Math.floor(maxVisiblePages / 2));
  let end = Math.min(totalPages.value, start + maxVisiblePages - 1);

  if (end - start + 1 < maxVisiblePages) {
    start = Math.max(1, end - maxVisiblePages + 1);
  }

  for (let i = start; i <= end; i++) {
    pages.push(i);
  }

  return pages;
});
</script>
