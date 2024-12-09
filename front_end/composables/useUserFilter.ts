import { getUsersQuery } from "~/queries/user";
import type { User } from "~/types/user";

export const useUserFilter = () => {
    const users = ref<User[]>([])
    const useFilter = useRecipeFilters()

    const fetchUsers = async () => {
        const variables = {
            username: useFilter.activeFilters.value.creator ?  `%${useFilter.activeFilters.value.creator}%` : ""
        }

        try {
          const { result } = useQuery(getUsersQuery, variables);
          console.log(result)
          watch(
            result,
            (newResult) => {
              if (newResult?.recipes) {
                users.value = newResult.recipes;
              }
            },
            { immediate: true }
          );
        } catch (error) {
          console.log(error)
        }
      }
      return {
        users,
        fetchUsers
      }
}