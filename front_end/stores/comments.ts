import { getCommentQuery } from "~/queries/recipe";


export const useCommentsStore = defineStore("comments", () => {
    const comments = ref([])

    const setComments = (data?: any) => (comments.value = data);
    
    const fetchComments = async (recipeId: number, userId: number) => {
      try {
        const { data, error }: any = await useAsyncQuery(getCommentQuery, {
          recipeId: recipeId,  userId: userId
        });
        
        console.log(data, error, "ssssssssssssssssssssssssss")
        if (!data.value) {
          throw new Error("no value returned");
        }
  
        watch(
          data,
          (newdata) => {
            if (newdata?.comments) {
              setComments(newdata.comments)
            }
          },
          { immediate: true }
        );
  
      } catch (error) {
        console.log(error);
        setComments([]);
      }
    }
  
    return {
      comments,
      setComments,
      fetchComments
    };
  });