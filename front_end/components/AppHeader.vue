<template>
  <header class="bg-white dark:bg-gray-800 shadow-sm">
    <nav class="container mx-auto px-4 py-4">
      <div class="flex items-center justify-between">
        <NuxtLink to="/" class="text-2xl font-bold text-primary-600 dark:text-primary-400">
          Recipe Haven
        </NuxtLink>
        
        <div class="flex items-center space-x-6">
          <SearchBar />
          <NavLinks />
          <ThemeToggle />
          <AuthSection />
          <UserMenu v-if="userStore.isAuthenticated()"/>
        </div>
      </div>
    </nav>
  </header>
</template>

<script setup lang="ts">
import SearchBar from '~/components/navigation/SearchBar.vue'
import NavLinks from '~/components/navigation/NavLinks.vue'
import AuthSection from '~/components/navigation/AuthSection.vue'
import UserMenu from '~/components/navigation/UserMenu.vue'

const userStore = useUserStore()
</script>