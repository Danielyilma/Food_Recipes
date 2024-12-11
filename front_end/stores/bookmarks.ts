import { getBookmarkedrecipes } from "~/queries/bookmarks";


export const useBookmarksStore = defineStore("bookmarks", () => {
    const bookmarks = ref<any>([])

    const setBookmarks = (data?: any) => (bookmarks.value = data);
    const fetchBookmarks = async (userId: number) => {
        try {
            const { data, error }: any = await useAsyncQuery(getBookmarkedrecipes, {userId})    
            watch(
                data,
                (newdata) => {
                  if (newdata?.bookmarks) {
                    console.log(newdata?.bookmarks)
                    setBookmarks(newdata.bookmarks);
                  }
                },
                { immediate: true }
              );
            
        } catch (error) {
            console.log(error)

        }
    }

    return {
        bookmarks,
        fetchBookmarks
    }
});