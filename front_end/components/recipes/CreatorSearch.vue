<template>
  <div class="relative">
    <input
      type="text"
      v-model="useFilter.activeFilters.value.creator"
      @input="useUsers.fetchUsers"
      @focus="showDropdown = true"
      class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-primary-500 focus:border-transparent"
      placeholder="Search creators..."
    />

    <!-- Dropdown -->
    <!-- <div
      v-if="showDropdown && useUsers.users.value.length > 0"
      class="absolute z-10 w-full mt-1 bg-white dark:bg-gray-700 rounded-md shadow-lg max-h-60 overflow-auto"
    >
      <div class="py-1">
        <button
          v-for="creator in useUsers.users.value"
          :key="creator.id"
          class="w-full px-4 py-2 text-left flex items-center space-x-3 hover:bg-gray-100 dark:hover:bg-gray-600"
        >
          <img
            :src="creator.profile_image"
            :alt="creator.username"
            class="w-8 h-8 rounded-full"
          />
          <span class="text-gray-900 dark:text-gray-100">{{
            creator.username
          }}</span>
        </button>
      </div>
    </div> -->
  </div>
</template>

<script setup lang="ts">

const useFilter = useRecipeFilters()
const useUsers = useUserFilter()


interface Creator {
  id: number;
  name: string;
  avatar: string;
  recipeCount: number;
}

const props = defineProps<{
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

// const searchQuery = ref("");
const showDropdown = ref(false);

// Sample creators data with avatars and recipe counts
// const creators: Creator[] = [
//   {
//     id: 1,
//     name: "John Doe",
//     avatar: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e",
//     recipeCount: 15,
//   },
//   {
//     id: 2,
//     name: "Jane Smith",
//     avatar: "https://images.unsplash.com/photo-1438761681033-6461ffad8d80",
//     recipeCount: 23,
//   },
//   {
//     id: 3,
//     name: "Mike Johnson",
//     avatar: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e",
//     recipeCount: 8,
//   },
// ];

// const filteredCreators = computed(() => {
//   if (!searchQuery.value) {
//     return creators.sort((a, b) => b.recipeCount - a.recipeCount);
//   }

//   const query = searchQuery.value.toLowerCase();
//   return creators
//     .filter((creator) => creator.name.toLowerCase().includes(query))
//     .sort((a, b) => b.recipeCount - a.recipeCount);
// });

const handleSearch = () => {
  showDropdown.value = true;
};

// const selectCreator = (creator: Creator) => {
//   searchQuery.value = creator.name;
//   emit("update:modelValue", creator.name);
//   showDropdown.value = false;
// };

// Close dropdown when clicking outside
//   onClickOutside(
//     computed(() => document.querySelector('.creator-search')),
//     () => {
//       showDropdown.value = false
//     }
//   )

// Watch for external value changes
// watch(
//   () => props.modelValue,
//   (newValue) => {
//     if (!newValue) {
//       searchQuery.value = "";
//     }
//   }
// );
</script>
