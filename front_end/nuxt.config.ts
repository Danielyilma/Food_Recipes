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
    "@pinia/nuxt",
    "@nuxtjs/color-mode",
  ],
  colorMode: {
    classSuffix: "",
    preference: "system",
    fallback: "light",
  },
  apollo: {
    authType: "Bearer",
    authHeader: "Authorization",
    clients: {
      default: {
        httpEndpoint: process.env.GRAPHQL_API || "",
      },
    },
  },
  pinia: {
    storesDirs: ["./stores/**", "./custom-folder/stores/**"],
  },
  build: {
    transpile: ["pinia-plugin-persistedstate"],
  },
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
});
