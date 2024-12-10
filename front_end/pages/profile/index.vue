<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-6xl mx-auto">
      <ProfileHeader 
        :user="userStore.user" 
        @edit="isEditModalOpen = true"
      />
      
      <div class="mt-8 grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Sidebar -->
        <div class="lg:col-span-1">
          <ProfileSidebar :user="userStore.user" />
        </div>
        
        <!-- Main Content -->
        <div class="lg:col-span-2">
          <ProfileRecipes :recipes="userStore.userRecipes" @update:recipes="userStore.userRecipes = $event" />
        </div>
      </div>

      <!-- Edit Profile Modal -->
      <ProfileEditModal
        :is-open="isEditModalOpen"
        :user="userStore.user"
        @close="isEditModalOpen = false"
        @update="handleProfileUpdate"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { User } from '~/types/user'
import ProfileEditModal from '~/components/profile/ProfileEditModal.vue';

const isEditModalOpen = ref(false)

definePageMeta({
    middleware: "authentication",
});

const userStore = useUserStore()

userStore.fetchUserRecipes(userStore.user.id)

// Mock data - In a real app, this would come from an API
// const user = ref<User>({
//   id: 1,
//   username: 'johndoe',
//   name: 'John Doe',
//   email: 'john@example.com',
//   avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e',
//   bio: 'Passionate food lover and home chef. I love creating and sharing delicious recipes!',
//   location: 'New York, USA',
//   joinedDate: '2024-01-15',
//   recipeCount: 15,
//   followersCount: 250,
//   followingCount: 120
// })

// const userRecipes = ref<Recipe[]>([
//   {
//     id: 1,
//     title: 'Classic Pancakes',
//     description: 'Fluffy and delicious pancakes perfect for breakfast',
//     image: 'https://images.unsplash.com/photo-1528207776546-365bb710ee93',
//     category: 'Breakfast',
//     prepTime: '20 mins',
//     createdAt: '2024-03-01',
//     likes: 42
//   }
// ])

const handleProfileUpdate = (updatedData: Partial<User>) => {
  // In a real app, this would make an API call
  Object.assign(user.value, updatedData)
}
</script>
