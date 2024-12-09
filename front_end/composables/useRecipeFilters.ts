import { getPrepTimeMinutes } from "~/utils/time";
import type { Recipe, Filters } from "~/types/recipe";
import { getFilteredRecipesQuery } from "~/queries/recipe";

export const useRecipeFilters = () => {
  const useRecipe = useRecipesStore();
  const searchQuery = ref("");
  const activeFilters = ref<Filters>({
    categories: [],
    prepTime: "",
    creator: "",
  });

  const fetchFilteredRecipes = async () => {
    try {
      const isEmpty = activeFilters.value.categories.length === 0 &&
      !activeFilters.value.creator && !searchQuery.value

      console.log(searchQuery.value, isEmpty, "te")
      if (isEmpty) {
        useRecipe.fetchRecipes(1)
        return;
      }

      const variables = {
        categories: activeFilters.value.categories || "",
        creator: activeFilters.value.creator ? `%${activeFilters.value.creator}%` : "",
        search: searchQuery.value ? "%" + searchQuery.value + "%" : "",
      };

      variables.search = variables.search === "%%" ? "" : variables.search;
      console.log(variables, "varibles");
      const { result, error }: any = useQuery(getFilteredRecipesQuery, variables);
      console.log(result, error)
      watch(
        result,
        (newResult) => {
          if (newResult?.recipes) {
            console.log(newResult.recipes, "tes")
            useRecipe.setRecipes(newResult.recipes);
          }
        },
        { immediate: true }
      );
    } catch (error) {
      console.log(error);
    }
  };

  const clearFilters = () => {
    activeFilters.value.categories = [],
    activeFilters.value.prepTime = "";
    activeFilters.value.creator = "";
    useRecipe.fetchRecipes(1)
  };

  watch([activeFilters], fetchFilteredRecipes, { deep: true });

  return {
    searchQuery,
    activeFilters,
    fetchFilteredRecipes,
    clearFilters
  };
};

// Apply search filter
// if (searchQuery.value) {
//   const query = searchQuery.value.toLowerCase();
//   filtered = filtered.filter(
//     (recipe) =>
//       recipe.title.toLowerCase().includes(query) ||
//       recipe.description.toLowerCase().includes(query)
//   );
// }

// // Apply category filter
// if (activeFilters.value.categories.length > 0) {
//   filtered = filtered.filter((recipe) =>
//     activeFilters.value.categories.includes(recipe.category)
//   );
// }

// // Apply prep time filter
// if (activeFilters.value.prepTime) {
//   filtered = filtered.filter((recipe) => {
//     const minutes = getPrepTimeMinutes(recipe.prepTime);
//     switch (activeFilters.value.prepTime) {
//       case "quick":
//         return minutes < 30;
//       case "medium":
//         return minutes >= 30 && minutes <= 60;
//       case "long":
//         return minutes > 60;
//       default:
//         return true;
//     }
//   });
// }

// // Apply creator filter
// if (activeFilters.value.creator) {
//   filtered = filtered.filter(
//     (recipe) => recipe.author.name === activeFilters.value.creator
//   );
// }

// watch(
//   filteredRecipes,
//   (newRecipes) => {
//     currentPage.value = 1; // Reset to first page when filters change
//   },
//   { deep: true }
// );
