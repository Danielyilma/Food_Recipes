// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      title: "Recipe Haven",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
      ],
    },
  },
  modules: [
    "@nuxtjs/apollo",
    "@nuxtjs/tailwindcss",
    // "@pinia/nuxt",
    "@nuxtjs/color-mode",
  ],
  colorMode: {
    classSuffix: "",
    preference: "system",
    fallback: "light",
  },
  apollo: {
    clients: {
      default: {
        httpEndpoint: "https://spacex-production.up.railway.app",
      },
    },
  },
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
});
