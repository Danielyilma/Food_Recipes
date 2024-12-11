import { useUserStore } from "~/stores/user";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const userStore = useUserStore();

  if (userStore.isAuthenticated() && ["/login", "/signup"].includes(to.path)) {
    return await navigateTo("/");
  }

  console.log("sssssssssssssssssssssssssssssss")
  if (
    !userStore.isAuthenticated() &&
    !["/login", "/signup"].includes(to.path)
  ) {
    return await navigateTo("/login");
  }
});
