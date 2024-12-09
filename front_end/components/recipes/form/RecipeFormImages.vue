<template>
    <div class="space-y-6">
      <div>
        <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Upload Recipe Images</h2>
        <div class="grid grid-cols-2 gap-4">
          <div
            v-for="(image, index) in images"
            :key="index"
            class="relative aspect-video rounded-lg overflow-hidden group"
          >
            <img :src="image" :alt="`Recipe image ${index + 1}`" class="w-full h-full object-cover">
            <button
              @click="removeImage(index)"
              class="absolute top-2 right-2 p-1 rounded-full bg-red-600 text-white opacity-0 group-hover:opacity-100 transition-opacity"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
            <button
              v-if="image !== featuredImage"
              @click="setFeaturedImage(image)"
              class="absolute bottom-2 right-2 px-2 py-1 text-xs rounded bg-primary-600 text-white opacity-0 group-hover:opacity-100 transition-opacity"
            >
              Set as Featured
            </button>
            <div
              v-else
              class="absolute bottom-2 right-2 px-2 py-1 text-xs rounded bg-green-600 text-white"
            >
              Featured Image
            </div>
          </div>
          
          <label
            class="border-2 border-dashed border-gray-300 dark:border-gray-600 rounded-lg p-4 flex flex-col items-center justify-center cursor-pointer hover:border-primary-500 dark:hover:border-primary-400"
          >
            <input
              type="file"
              multiple
              accept="image/*"
              class="hidden"
              @change="handleImageUpload"
            >
            <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            <span class="mt-2 text-sm text-gray-600 dark:text-gray-400">Add Images</span>
          </label>
        </div>
      </div>
      
      <div class="flex justify-end">
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
  const props = defineProps<{
    images: string[]
    featuredImage: string
  }>()
  
  const emit = defineEmits<{
    (e: 'update:images', images: string[]): void
    (e: 'update:featuredImage', image: string): void
    (e: 'next'): void
  }>()
  
  const isValid = computed(() => props.images.length > 0 && props.featuredImage)
  
  const handleImageUpload = (event: Event) => {
    const files = (event.target as HTMLInputElement).files
    if (!files) return
    console.log(files)
    // In a real app, this would upload to a server
    Array.from(files).forEach(file => {
      const reader = new FileReader()
      reader.onload = (e) => {
        const newImages = [...props.images, e.target?.result as string]
        emit('update:images', newImages)
        if (!props.featuredImage) {
          emit('update:featuredImage', e.target?.result as string)
        }
      }
      reader.readAsDataURL(file)
    })
  }
  
  const removeImage = (index: number) => {
    const newImages = props.images.filter((_, i) => i !== index)
    emit('update:images', newImages)
    if (props.images[index] === props.featuredImage) {
      emit('update:featuredImage', newImages[0] || '')
    }
  }
  
  const setFeaturedImage = (image: string) => {
    emit('update:featuredImage', image)
  }
  </script>