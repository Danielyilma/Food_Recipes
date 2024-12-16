<template>
  <div class="space-y-6">
    <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
      Comments
    </h2>

    <!-- New Comment Form -->
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <textarea
        v-model="newComment"
        rows="3"
        class="w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-primary-500 focus:border-transparent"
        placeholder="Share your thoughts..."
      ></textarea>

      <button
        type="submit"
        class="px-4 py-2 bg-primary-600 hover:bg-primary-700 text-white rounded-lg font-medium"
      >
        Post Comment
      </button>
    </form>

    <!-- Comments List -->
    <div class="space-y-6">
      <div v-for="comment in useComments?.comments" :key="comment.id" class="flex space-x-4">
        <img
          :src="config.public.imageDomainPath + comment?.user?.profile_image"
          :alt="comment?.user?.username"
          class="w-10 h-10 rounded-full"
        />

        <div class="flex-1">
          <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4">
            <div class="flex items-center justify-between mb-2">
              <span class="font-medium text-gray-900 dark:text-white">{{
                comment?.user?.username
              }}</span>
              <span class="text-sm text-gray-500 dark:text-gray-400">{{
                formatDate(comment?.created_at)
              }}</span>
            </div>
            <p class="text-gray-700 dark:text-gray-300">
              {{ comment?.content }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { insertRecipeCommentQuery } from "~/queries/recipe";
import { useCommentsStore } from "~/stores/comments";
import type { Comment } from "~/types/recipe";

const props = defineProps<{
  recipe_id: number;
}>();

const userStore = useUserStore()
const useComments = useCommentsStore();
const config = useRuntimeConfig()

await useComments.fetchComments(props.recipe_id)
const newComment = ref("");

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
  });
};

const handleSubmit = async () => {
  if (newComment.value.trim()) {
    const variable: any = {
      recipeId: props.recipe_id,
      content: newComment.value,
      userId: userStore.user.id
    }
    console.log(variable)
    try {
      const { mutate } = useMutation<any>(insertRecipeCommentQuery, variable) 
      const {data}: any = await mutate(variable)

      if (data?.insert_comments?.returning?.[0]) {
        const addedComment: any = data.insert_comments.returning[0];
        useComments.comments.push(addedComment);
      }
    } catch (error) {
      console.log(error)
    }

    newComment.value = "";
  }
};
</script>
