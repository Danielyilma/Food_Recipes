import { getCategorisQuery } from "~/queries/categories";
import type { Category } from "~/types/recipe";


export const useCategoryStore = defineStore("categories", () => {
    const categories = ref<Category[]>([])

    const setCategories = (data?: any) => (categories.value = data);
    const fetchCategories = async () => {
        try {
            const { data }: any = await useAsyncQuery(getCategorisQuery, {})

            watch(
                data,
                (newdata) => {
                  if (newdata?.categories) {
                    setCategories(newdata.categories);
                  }
                },
                { immediate: true }
              );
            
        } catch (error) {
            console.log(error)
        }
    }

    return {
        categories,
        fetchCategories
    }
});
