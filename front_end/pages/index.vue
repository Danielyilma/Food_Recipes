<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative bg-center bg-cover h-[600px]"
      style="
        background-image: url('https://images.unsplash.com/photo-1498837167922-ddd27525d352?ixlib=rb-1.2.1&auto=format&fit=crop&w=2850&q=80');
      "
    >
      <div class="absolute inset-0 bg-black bg-opacity-50"></div>
      <div class="relative container mx-auto px-4 h-full flex items-center">
        <div class="max-w-2xl">
          <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">
            Discover & Share Amazing Recipes
          </h1>
          <p class="text-xl text-gray-200 mb-8">
            Join our community of food lovers and share your favorite recipes
            with the world
          </p>
          <div class="space-x-4">
            <NuxtLink
              to="/recipes"
              class="bg-primary-600 hover:bg-primary-700 text-white px-6 py-3 rounded-lg font-medium"
            >
              Explore Recipes
            </NuxtLink>
            <NuxtLink
              to="/share-recipe"
              class="bg-white hover:bg-gray-100 text-primary-600 px-6 py-3 rounded-lg font-medium"
            >
              Share Recipe
            </NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- Categories Section -->
    <section class="py-16">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-12 dark:text-white">
          Popular Categories
        </h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
          <div v-for="category in categories" :key="category.id" class="group">
            <NuxtLink :to="'/categories/' + category.slug" class="block">
              <div class="relative rounded-lg overflow-hidden">
                <img
                  :src="category.image"
                  :alt="category.name"
                  class="w-full h-48 object-cover transition duration-300 group-hover:scale-110"
                />
                <div
                  class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"
                ></div>
                <h3
                  class="absolute bottom-4 left-4 text-white text-xl font-semibold"
                >
                  {{ category.name }}
                </h3>
              </div>
            </NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Recipes Section -->
    <section class="py-16 bg-gray-100 dark:bg-gray-800">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-12 dark:text-white">
          Featured Recipes
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div
            v-for="recipe in featuredRecipes"
            :key="recipe.id"
            class="bg-white dark:bg-gray-700 rounded-lg overflow-hidden shadow-lg"
          >
            <img
              :src="recipe.image"
              :alt="recipe.title"
              class="w-full h-48 object-cover"
            />
            <div class="p-6">
              <h3 class="text-xl font-semibold mb-2 dark:text-white">
                {{ recipe.title }}
              </h3>
              <p class="text-gray-600 dark:text-gray-300 mb-4">
                {{ recipe.description }}
              </p>
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <img
                    :src="recipe.author.avatar"
                    :alt="recipe.author.name"
                    class="w-8 h-8 rounded-full"
                  />
                  <span class="ml-2 text-sm text-gray-600 dark:text-gray-300">{{
                    recipe.author.name
                  }}</span>
                </div>
                <NuxtLink
                  :to="'/recipes/' + recipe.id"
                  class="text-primary-600 hover:text-primary-700 dark:text-primary-400 dark:hover:text-primary-500 font-medium"
                >
                  View Recipe
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
const categories = [
  {
    id: 1,
    name: "Breakfast",
    slug: "breakfast",
    image:
      "https://images.unsplash.com/photo-1533089860892-a7c6f0a88666?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
  },
  {
    id: 2,
    name: "Main Course",
    slug: "main-course",
    image:
      "https://images.unsplash.com/photo-1547928576-965be7f5f6a2?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
  },
  {
    id: 3,
    name: "Desserts",
    slug: "desserts",
    image:
      "https://images.unsplash.com/photo-1488477181946-6428a0291777?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
  },
  {
    id: 4,
    name: "Healthy",
    slug: "healthy",
    image:
      "https://images.unsplash.com/photo-1498837167922-ddd27525d352?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
  },
];

const featuredRecipes = [
  {
    id: 1,
    title: "Classic Pancakes",
    description: "Fluffy and delicious pancakes perfect for breakfast",
    image:
      "https://images.unsplash.com/photo-1528207776546-365bb710ee93?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
    author: {
      name: "John Doe",
      avatar:
        "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
  },
  {
    id: 2,
    title: "Grilled Salmon",
    description: "Healthy and tasty grilled salmon with vegetables",
    image:
      "https://images.unsplash.com/photo-1467003909585-2f8a72700288?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
    author: {
      name: "Jane Smith",
      avatar:
        "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
  },
  {
    id: 3,
    title: "Chocolate Cake",
    description: "Rich and moist chocolate cake for special occasions",
    image:
      "https://images.unsplash.com/photo-1578985545062-69928b1d9587?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
    author: {
      name: "Mike Johnson",
      avatar:
        "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
  },
];
</script>
