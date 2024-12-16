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
        httpEndpoint: process.env.NUXT_GRAPHQL_API || "",
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
  runtimeConfig: {
    public: {
      fileUploadApi: process.env.NUXT_FILE_UPLOAD_API || '',
      imageDomainPath: process.env.NUXT_IMAGE_DOMAIN_PATH || '',
      chapaApiKey: process.env.NUXT_CHAPA_API_KEY || '',
      chapaInitApi: process.env.NUXT_CHAPA_INIT_API || ''
    }
  }
});
