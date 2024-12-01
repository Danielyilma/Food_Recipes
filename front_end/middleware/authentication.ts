import { useUserStore } from "~/stores/user";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const userStore = useUserStore();

  if (userStore.isAuthenticated() && ["/login", "/signup"].includes(to.path)) {
    return await navigateTo("/");
  }

  if (!userStore.isAuthenticated() && to.path !== "/login") {
    return await navigateTo("/login");
  }
});
