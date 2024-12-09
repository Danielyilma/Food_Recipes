
export const sendMutation = async (query: any, variable: any) => {
    try {
        const { mutate } = useMutation(query, variable)
        const { data }: any = await mutate(variable)
        return data
    } catch (error) {
        console.log(error)
    }
}

export const sendQuery = async (query: any, variables: any) => {
    console.log("query:", query)
    console.log("variable:", variables)
    try {
      const { data, error } = await useAsyncQuery(query, variables);
  
      if (error) {
        console.error("GraphQL Query Error:", error);
        return null;
      }
  
      console.log("GraphQL Query Result:", data);
      return data;
    } catch (err) {
      console.error("Unexpected Error:", err);
      return null;
    }
  };
  