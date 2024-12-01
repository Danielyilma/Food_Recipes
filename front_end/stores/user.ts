import { getUserQuery, loginQuery } from "~/queries/user";
import { jwtDecode } from "jwt-decode";
import type { CustomJwtPayload } from "~/types";

export const useUserStore = defineStore("user", () => {
  const user = ref();
  const token = useCookie("recipe_token_864450f1-21a6-4fbf-be46-7ee6789af902", {
    maxAge: 60 * 60,
  });

  const setToken = (data?: any) => {
    console.log(data);
    token.value = data;
  };
  const setUser = (data?: any) => {
    user.value = data;
  };

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
        const { data }: any = await useAsyncQuery(getUserQuery, variables);
        setUser(data.value.users[0]);
        // fetch user data
      } catch (error) {
        setUser();
        console.log(error);
      }
    }
  };

  const logout = () => {
    setToken();
    setToken();
  };

  return {
    user,
    token,
    setToken,
    setUser,
    login,
    fetchUser,
    logout,
    isAuthenticated,
  };
});
