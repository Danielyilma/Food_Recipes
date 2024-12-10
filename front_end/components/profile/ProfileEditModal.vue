<template>
    <TransitionRoot appear :show="isOpen" as="template">
      <Dialog as="div" @close="closeModal" class="relative z-50">
        <TransitionChild
          enter="duration-300 ease-out"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div class="fixed inset-0 bg-black/25 dark:bg-black/40" />
        </TransitionChild>
  
        <div class="fixed inset-0 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4">
            <TransitionChild
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <DialogPanel class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white dark:bg-gray-800 p-6 shadow-xl transition-all">
                <DialogTitle class="text-lg font-medium leading-6 text-gray-900 dark:text-white mb-4">
                  Edit Profile
                </DialogTitle>
  
                <form @submit.prevent="handleSubmit" class="space-y-4">
                  <div class="flex items-center space-x-4">
                    <div class="relative">
                      <img 
                        :src="formData.image_url"
                        class="w-20 h-20 rounded-full object-cover"
                        alt="Profile avatar"
                      >
                      <button
                        type="button"
                        @click="triggerFileInput"
                        class="absolute bottom-0 right-0 bg-primary-600 text-white p-1.5 rounded-full hover:bg-primary-700"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                        </svg>
                      </button>
                      <input
                        ref="fileInput"
                        type="file"
                        accept="image/*"
                        class="hidden"
                        @change="handleAvatarChange"
                      >
                    </div>
                    <div>
                      <h4 class="text-sm font-medium text-gray-900 dark:text-white">Profile Photo</h4>
                      <p class="text-sm text-gray-500 dark:text-gray-400">Click the camera icon to change</p>
                    </div>
                  </div>
  
                  <div>
                    <FormInput
                      id="name"
                      label="Name"
                      v-model="formData.name"
                      placeholder="Enter your name"
                      required
                    />
                  </div>
  
                  <div>
                    <FormInput
                      id="username"
                      label="Username"
                      v-model="formData.username"
                      placeholder="Choose a username"
                      required
                    />
                  </div>
  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                      Bio
                    </label>
                    <textarea
                      v-model="formData.bio"
                      rows="3"
                      class="w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm focus:border-primary-500 focus:ring-primary-500 dark:bg-gray-700 dark:text-white"
                      placeholder="Tell us about yourself"
                    ></textarea>
                  </div>
  
                  <!-- <div>
                    <FormInput
                      id="location"
                      label="Location"
                      v-model="formData.location"
                      placeholder="Your location"
                    />
                  </div> -->
  
                  <div class="flex justify-end space-x-3 mt-6">
                    <button
                      type="button"
                      @click="closeModal"
                      class="px-4 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-md"
                    >
                      Cancel
                    </button>
                    <button
                      type="submit"
                      class="px-4 py-2 text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 rounded-md"
                    >
                      Save Changes
                    </button>
                  </div>
                </form>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import {
    Dialog,
    DialogPanel,
    DialogTitle,
    TransitionChild,
    TransitionRoot,
  } from '@headlessui/vue'
  import type { User } from '~/types/user'
import { updateUserQuery } from '~/queries/user';
  
  const props = defineProps<{
    isOpen: boolean
    user: User
  }>()
  const userStore = useUserStore()
  const config = useRuntimeConfig()
  
  const emit = defineEmits<{
    (e: 'close'): void
    (e: 'update', data: Partial<User>): void
  }>()
  
  const fileInput = ref<HTMLInputElement | null>(null)
  const formData = ref({
    username: props.user.username,
    bio: props.user.bio,
    image_url: config.public.imageDomainPath + props.user.profile_image,
    image: ""
  })
  
  const triggerFileInput = () => {
    fileInput.value?.click()
  }
  
  const handleAvatarChange = (event: Event) => {
    const file = (event.target as HTMLInputElement).files?.[0]
    if (file) {
      const reader = new FileReader()
      reader.onload = (e) => {
        formData.value.image = e.target?.result as string
        formData.value.image_url = e.target?.result as string
      }
      reader.readAsDataURL(file)
    }
  }
  
  const closeModal = () => {
    emit('close')
  }
  
  const handleSubmit = async() => {
    const formsData = new FormData()
    formsData.append("file", formData.value.image);

    const uploadResponse = await fetch(config.public.fileUploadApi + "single-upload", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${userStore.token}`
        },
        body: formsData
      });

      // if (!uploadResponse.ok) {
      //   // return;
      // }
      
      const { url } = await uploadResponse.json();
      formData.value.image = url
      await sendMutation(updateUserQuery, {id: userStore.user.id, bio: formData.value.bio, profileImage: url || formData.value.image_url})
      userStore.user.profile_image = url
      userStore.user.bio = formData.value.bio
    closeModal()
  }
  </script>