<template>
  <ClientOnly>
    <AuthCard title="Create an Account">
      <template #subtitle>
        Already have an account?
        <NuxtLink
          to="/login"
          class="font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400"
        >
          Sign in
        </NuxtLink>
      </template>

      <form @submit.prevent="handleSignup" class="mt-8 space-y-6">
        <div class="space-y-4">
          <FormInput
            id="username"
            label="Username"
            type="text"
            v-model="SignupData.username"
            placeholder="Choose a username"
          />

          <FormInput
            id="email"
            label="Email address"
            type="email"
            v-model="SignupData.email"
            placeholder="Enter your email"
          />

          <FormInput
            id="phone_number"
            label="Phone number"
            type="text"
            v-model="SignupData.phone_number"
            placeholder="Your phone number"
          />

          <FormInput
            id="password"
            label="Password"
            type="password"
            v-model="SignupData.password"
            placeholder="Create a password"
          />

          <FormInput
            id="confirmPassword"
            label="Confirm Password"
            type="password"
            v-model="SignupData.confirmPassword"
            placeholder="Confirm your password"
          />

          <div class="flex items-center">
            <input
              id="terms"
              type="checkbox"
              v-model="acceptTerms"
              class="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded"
            />
            <label
              for="terms"
              class="ml-2 block text-sm text-gray-700 dark:text-gray-300"
            >
              I agree to the
              <NuxtLink
                to="/terms"
                class="font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400"
              >
                Terms of Service
              </NuxtLink>
              and
              <NuxtLink
                to="/privacy"
                class="font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400"
              >
                Privacy Policy
              </NuxtLink>
            </label>
          </div>
        </div>

        <button
          type="submit"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
          @click.prevent="signup"
        >
          Create Account
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

        <SocialLogin />
      </form>
    </AuthCard>
  </ClientOnly>
</template>

<script setup>
import FormInput from "~/components/ui/FormInput.vue";
import { signupQuery } from "~/queries/user";

definePageMeta({
  layout: "empty",
  middleware: "authentication",
});

const SignupData = ref({
  username: "",
  email: "",
  password: "",
  phone_number: "",
  confirmPassword: "",
});

const signup = async () => {
  const data = {
    username: SignupData.value.username,
    password: SignupData.value.password,
    email: SignupData.value.email,
    phone_number: SignupData.value.phone_number,
  };

  try {
    const { mutate } = useMutation(signupQuery, data);
    await mutate(data);
    return await navigateTo("/login", { replace: true });
  } catch (error) {
    console.log(error);
  }
};
</script>
