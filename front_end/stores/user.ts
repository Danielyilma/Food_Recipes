import { getUserQuery, getUsersQuery, loginQuery } from "~/queries/user";
import { jwtDecode } from "jwt-decode";
import type { CustomJwtPayload } from "~/types/user";
import { getUserRecipeQuery } from "~/queries/recipe";

export const useUserStore = defineStore("user", () => {
  const user = ref();
  const recipeIsliked = ref()
  const userRecipes = ref([])
  const token = useCookie("recipe_token_864450f1-21a6-4fbf-be46-7ee6789af902", {
    maxAge: 60 * 60,
  });

  const setToken = (data?: any) => {
    token.value = data;
  };
  const setUser = (data?: any) => {
    user.value = data;
  };
  const setUserRecipes = (data?: any) => {
    userRecipes.value = data
  }

  const isAuthenticated = () => !!user.value;

  const login = async (loginData: any) => {
    try {
      const { mutate } = useMutation<any>(loginQuery, loginData);
      const { data }: any = await mutate(loginData);
      setToken(data.login.token);
      await fetchUser();
    } catch (error) {
      setToken();
      setUser();
    }
  };

  const fetchUser = async () => {
    if (token.value) {
      try {
        const tokenItem = jwtDecode<CustomJwtPayload>(token.value);
        const variables = {
          id: tokenItem.user_id,
        };

        const { data, error }: any = await useAsyncQuery(getUserQuery, variables);
        console.log(error)
        setUser(data.value.users[0]);
        // fetch user data
      } catch (error) {
        setUser();
        console.log(error);
      }
    }
  };

  const fetchUserRecipes = async (id: number) => {
    try {
      const { data, error }:any = useAsyncQuery(getUserRecipeQuery, {id: id})
      
      setUserRecipes(data.value.recipes)
    } catch (error) {
        console.log(error)
    }
  }


  const logout = () => {
    setToken();
    setToken();
  };

  return {
    user,
    userRecipes,
    token,
    setToken,
    setUser,
    login,
    fetchUser,
    fetchUserRecipes,
    logout,
    isAuthenticated,
  };
});
