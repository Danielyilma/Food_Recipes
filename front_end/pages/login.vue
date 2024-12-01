<template>
  <ClientOnly>
    <AuthCard title="Welcome Back">
      <template #subtitle>
        Don't have an account?
        <NuxtLink
          to="/signup"
          class="font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400"
        >
          Sign up
        </NuxtLink>
      </template>

      <form @submit.prevent="handleLogin" class="mt-8 space-y-6">
        <div class="space-y-4">
          <FormInput
            id="email"
            label="Email address"
            type="email"
            v-model="email"
            placeholder="Enter your email"
          />

          <FormInput
            id="password"
            label="Password"
            type="password"
            v-model="password"
            placeholder="Enter your password"
          />

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="remember-me"
                type="checkbox"
                v-model="rememberMe"
                class="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded"
              />
              <label
                for="remember-me"
                class="ml-2 block text-sm text-gray-700 dark:text-gray-300"
              >
                Remember me
              </label>
            </div>

            <NuxtLink
              to="/forgot-password"
              class="text-sm font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400"
            >
              Forgot password?
            </NuxtLink>
          </div>
        </div>

        <button
          @click.prevent="signin"
          type="submit"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          Sign in
        </button>

        <div class="relative my-6">
          <div class="absolute inset-0 flex items-center">
            <div
              class="w-full border-t border-gray-300 dark:border-gray-600"
            ></div>
          </div>
          <!-- <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-white dark:bg-gray-800 text-gray-500"
              >Or continue with</span
            >
          </div> -->
        </div>

        <!-- <SocialLogin /> -->
      </form>
    </AuthCard>
  </ClientOnly>
</template>

<script setup>
import FormInput from "~/components/ui/FormInput.vue";

const userStore = useUserStore();

const email = ref("");
const password = ref("");

const signin = async () => {
  await userStore.login({ email: email.value, password: password.value });
  email.value = password.value = "";
  return await navigateTo("/", { replace: true });
};

definePageMeta({
  layout: "empty",
  middleware: "authentication",
});
</script>
