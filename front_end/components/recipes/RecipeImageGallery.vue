<template>
    <div class="space-y-4">
      <!-- Featured Image -->
      <div class="relative h-96 rounded-lg overflow-hidden">
        <img 
          :src="config.public.imageDomainPath + currentImage" 
          :alt="title"
          class="absolute inset-0 w-full h-full object-cover"
        >
        
        <!-- Navigation Arrows -->
        <div class="absolute inset-0 flex items-center justify-between p-4">
          <button 
            @click="previousImage" 
            class="p-2 rounded-full bg-black/50 text-white hover:bg-black/70 transition-colors"
            :class="{ 'opacity-50 cursor-not-allowed': currentIndex === 0 }"
            :disabled="currentIndex === 0"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          
          <button 
            @click="nextImage" 
            class="p-2 rounded-full bg-black/50 text-white hover:bg-black/70 transition-colors"
            :class="{ 'opacity-50 cursor-not-allowed': currentIndex === images.length - 1 }"
            :disabled="currentIndex === images.length - 1"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
        
        <!-- Image Counter -->
        <div class="absolute bottom-4 right-4 px-3 py-1 rounded-full bg-black/50 text-white text-sm">
          {{ currentIndex + 1 }} / {{ images.length }}
        </div>
      </div>
      
      <!-- Thumbnail Strip -->
      <div class="flex space-x-2 overflow-x-auto pb-2">
        <button
          v-for="(image, index) in images"
          :key="index"
          @click="setCurrentImage(index)"
          class="flex-shrink-0 w-20 h-20 rounded-lg overflow-hidden focus:outline-none focus:ring-2 focus:ring-primary-500"
          :class="{ 'ring-2 ring-primary-500': currentIndex === index }"
        >
          <img 
            :src="config.public.imageDomainPath + image?.image_url" 
            :alt="`${title} - Image ${index + 1}`"
            class="w-full h-full object-cover"
          >
        </button>
      </div>
    </div>
  </template>
  
<script setup lang="ts">
  import type { RecipeImage } from '~/types/recipe';

  const config = useRuntimeConfig()
  const props = defineProps<{
    images: RecipeImage[]
    title: string
  }>()
  
  const currentIndex = ref(0)
  const currentImage = computed(() => props.images[currentIndex.value]?.image_url)
  
  const nextImage = () => {
    if (currentIndex.value < props.images.length - 1) {
      currentIndex.value++
    }
  }
  
  const previousImage = () => {
    if (currentIndex.value > 0) {
      currentIndex.value--
    }
  }
  
  const setCurrentImage = (index: number) => {
    currentIndex.value = index
  }
  </script>