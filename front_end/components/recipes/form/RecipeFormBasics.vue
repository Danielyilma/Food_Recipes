<template>
    <div class="space-y-6">
      <div>
        <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Basic Information</h2>
        
        <div class="space-y-4">
          <FormInput
            id="title"
            label="Recipe Title"
            v-model="title"
            placeholder="Enter recipe title"
            @input="$emit('update:title', $event.target.value)"
          />
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              Description
            </label>
            <textarea
              v-model="description"
              @input="$emit('update:description', $event.target.value)"
              rows="4"
              class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
              placeholder="Describe your recipe"
            ></textarea>
          </div>
        </div>
      </div>
      
      <div class="flex justify-between">
        <button
          @click="$emit('back')"
          class="px-4 py-2 border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700"
        >
          Back
        </button>
        <button
          @click="$emit('next')"
          :disabled="!isValid"
          class="px-4 py-2 bg-primary-600 text-white rounded-lg disabled:opacity-50"
        >
          Next Step
        </button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import FormInput from "~/components/ui/FormInput.vue"
  const props = defineProps<{
    title: string
    description: string
  }>()
  
  const title = ref(props.title)
  const description = ref(props.description)
  
  defineEmits<{
    (e: 'update:title', value: any): void
    (e: 'update:description', value: any): void
    (e: 'next'): void
    (e: 'back'): void
  }>()
  
  const isValid = computed(() => 
    title.value.trim().length > 0 && 
    description.value.trim().length > 0
  )
  </script>