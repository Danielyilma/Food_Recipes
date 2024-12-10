<template>
    <Menu as="div" class="relative">
      <MenuButton class="flex items-center space-x-2">
        <img 
          :src="config.public.imageDomainPath + userStore.user.profile_image" 
          :alt="userStore.user.username"
          class="w-8 h-8 rounded-full object-cover"
        >
        <span class="text-gray-700 dark:text-gray-200">{{ userStore.user.username }}</span>
      </MenuButton>
  
      <transition
        enter-active-class="transition duration-100 ease-out"
        enter-from-class="transform scale-95 opacity-0"
        enter-to-class="transform scale-100 opacity-100"
        leave-active-class="transition duration-75 ease-in"
        leave-from-class="transform scale-100 opacity-100"
        leave-to-class="transform scale-95 opacity-0"
      >
        <MenuItems class="absolute right-0 mt-2 w-48 origin-top-right rounded-md bg-white dark:bg-gray-700 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none z-10">
          <div class="py-1">
            <MenuItem v-slot="{ active }">
              <NuxtLink
                to="/profile"
                :class="[
                  active ? 'bg-gray-100 dark:bg-gray-600' : '',
                  'block px-4 py-2 text-sm text-gray-700 dark:text-gray-200'
                ]"
              >
                Profile
              </NuxtLink>
            </MenuItem>
            <MenuItem v-slot="{ active }">
              <button
                @click="handleLogout"
                :class="[
                  active ? 'bg-gray-100 dark:bg-gray-600' : '',
                  'block w-full text-left px-4 py-2 text-sm text-gray-700 dark:text-gray-200'
                ]"
              >
                Logout
              </button>
            </MenuItem>
          </div>
        </MenuItems>
      </transition>
    </Menu>
  </template>
  
  <script setup lang="ts">
  import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
  
  const userStore = useUserStore()
  const config = useRuntimeConfig()

  const handleLogout = async () => {
    // Implement logout functionality
    userStore.logout()
    useRouter().go(0)
  }
  </script>