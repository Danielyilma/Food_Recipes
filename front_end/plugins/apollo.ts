export default defineNuxtPlugin((nuxtApp) => {
  const userStore = useUserStore();

  nuxtApp.hook("apollo:auth", ({ client, token }) => {
    if (userStore.token) {
      token.value = userStore.token;
    }
  });
});
