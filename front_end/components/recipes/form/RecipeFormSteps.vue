<template>
    <div class="space-y-6">
      <div>
        <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Recipe Instructions</h2>
        
        <div class="space-y-6">
          <div
            v-for="(step, index) in steps"
            :key="index"
            class="space-y-4"
          >
            <div class="flex items-start gap-4">
              <div class="flex-shrink-0 w-8 h-8 rounded-full bg-primary-100 dark:bg-primary-800 flex items-center justify-center">
                <span class="text-primary-600 dark:text-primary-400 font-medium">{{ index + 1 }}</span>
              </div>
              <div class="flex-1">
                <textarea
                  v-model="step.description"
                  @input="updateSteps"
                  rows="3"
                  class="block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
                  :placeholder="`Step ${index + 1} instructions`"
                ></textarea>
              </div>
              <button
                @click="removeStep(index)"
                class="p-2 text-red-600 hover:text-red-700 dark:text-red-400"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
            
            <div class="ml-12">
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Step Image (optional)</label>
              <div class="flex items-center gap-4">
                <img
                  v-if="step.image"
                  :src="step?.image"
                  :alt="`Step ${index + 1}`"
                  class="w-24 h-24 object-cover rounded-lg"
                >
                <label class="flex-1 cursor-pointer">
                  <input
                    type="file"
                    accept="image/*"
                    class="hidden"
                    @change="(e) => handleStepImageUpload(e, index)"
                  >
                  <div class="border-2 border-dashed border-gray-300 dark:border-gray-600 rounded-lg p-4 flex flex-col items-center justify-center hover:border-primary-500 dark:hover:border-primary-400">
                    <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <span class="mt-2 text-sm text-gray-600 dark:text-gray-400">
                      {{ step.image ? 'Change Image' : 'Add Image' }}
                    </span>
                  </div>
                </label>
              </div>
            </div>
          </div>
          
          <button
            @click="addStep"
            class="flex items-center text-primary-600 hover:text-primary-700 dark:text-primary-400"
          >
            <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add Step
          </button>
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
  interface Step {
    description: string
    image?: string
  }
  
  const props = defineProps<{
    steps: Step[]
  }>()
  
  const emit = defineEmits<{
    (e: 'update:steps', value: Step[]): void
    (e: 'next'): void
    (e: 'back'): void
  }>()
  
  const addStep = () => {
    emit('update:steps', [
      ...props.steps,
      { description: '' }
    ])
  }
  
  const removeStep = (index: number) => {
    const newSteps = [...props.steps]
    newSteps.splice(index, 1)
    emit('update:steps', newSteps)
  }
  
  const updateSteps = () => {
    emit('update:steps', [...props.steps])
  }
  
  const handleStepImageUpload = (event: Event, index: number) => {
    const file = (event.target as HTMLInputElement).files?.[0]
    if (!file) return
  
    // In a real app, this would upload to a server
    const reader = new FileReader()
    reader.onload = (e) => {
      const newSteps = [...props.steps]
      newSteps[index] = {
        ...newSteps[index],
        image: e.target?.result as string
      }
      emit('update:steps', newSteps)
    }
    reader.readAsDataURL(file)
  }
  
  const isValid = computed(() => 
    props.steps.length > 0 &&
    props.steps.every(step => step.description.trim())
  )
  </script>