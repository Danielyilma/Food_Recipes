<template>
  <header class="bg-white dark:bg-gray-800 shadow-sm">
    <nav class="container mx-auto px-4 py-4">
      <div class="flex items-center justify-between">
        <NuxtLink
          to="/"
          class="text-2xl font-bold text-primary-600 dark:text-primary-400"
        >
          Recipe Haven
        </NuxtLink>

        <div class="flex items-center space-x-6">
          <div class="relative">
            <input
              type="search"
              placeholder="Search recipes..."
              class="w-64 px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-primary-500"
            />
          </div>

          <div class="flex items-center space-x-4">
            <NuxtLink
              to="/recipes"
              class="text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            >
              Recipes
            </NuxtLink>
            <NuxtLink
              to="/categories"
              class="text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            >
              Categories
            </NuxtLink>
            <ThemeToggle />
            <NuxtLink
              to="/login"
              class="text-gray-600 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            >
              Login
            </NuxtLink>
          </div>
        </div>
      </div>
    </nav>
  </header>
</template>
