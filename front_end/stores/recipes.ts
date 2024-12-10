import { getRecipesQuery, getRecipeQuery, getLikeCountQuery, getLikeCountandislikedQuery } from "~/queries/recipe";
import type { Recipe } from "~/types/recipe";

export const useRecipesStore = defineStore("recipes", () => {
  const recipes = ref<Recipe[]>([]);
  const recipe = ref({})
  const isLiked = ref()
  const recipeLikes = ref()
  const totalRecipes = ref();
  const currentPage = ref(1);
  const pageSize = ref(10);

  const setRecipes = (data?: any) => (recipes.value = data);
  const setTotalRecipe = (data?: any) => (totalRecipes.value = data);
  const setCurrentPage = (data?: any) => (currentPage.value = data);
  const setRecipe = (data?: any) => (recipe.value = data)
  const setRecipelikes = (data?: any) => (recipeLikes.value = data)
  const setisLiked = (data?: any) => (isLiked.value = data)

  const fetchRecipes = async (page: number) => {
    try {
      const limit = pageSize.value;
      const offset: number = (page - 1) * limit;
      const { data, error }: any = await useAsyncQuery(getRecipesQuery, {
        limit,
        offset,
      });

      if (!data.value) {
        throw new Error("no value returned");
      }
      setRecipes(data.value.recipes);
      setTotalRecipe(data.value.recipes_aggregate.aggregate.count);
      setCurrentPage(page);
    } catch (error) {
      console.log(error);
      setRecipes([]);
    }
  };

  const fetchRecipe = async (id: number) => {
    try {
      const { data, error }: any = await useAsyncQuery(getRecipeQuery, {
        id
      });
      
      console.log(data, error)
      if (!data.value) {
        throw new Error("no value returned");
      }

      watch(
        data,
        (newdata) => {
          if (newdata?.recipes_by_pk) {
            console.log(newdata, "in watch")
            setRecipe(newdata.recipes_by_pk)
          }
        },
        { immediate: true }
      );

    } catch (error) {
      console.log(error);
      setRecipes([]);
    }
  }

  const fetchRecipeLikes = async (user_id: number, recipe_id: number) => {
    try {
      const { data, error }: any = await useAsyncQuery(getLikeCountandislikedQuery, {
        recipeId: recipe_id, userId: user_id
      });
      
      console.log(data, error)
      if (!data.value) {
        throw new Error("no value returned");
      }

      watch(
        data,
        (newdata) => {
          if (newdata?.likes_aggregate) {
            setRecipelikes(newdata?.likes_aggregate?.aggregate?.count)
            setisLiked(newdata?.likes.length > 0)
            console.log(recipeLikes, isLiked, "check")
          }
        },
        { immediate: true }
      );

    } catch (error) {
      console.log(error);
      setRecipes([]);
    }
  }

  return {
    recipe,
    recipeLikes,
    isLiked,
    recipes,
    currentPage,
    fetchRecipe,
    fetchRecipes,
    fetchRecipeLikes,
    setRecipes,
  };
});
