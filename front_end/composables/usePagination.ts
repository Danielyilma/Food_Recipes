export function usePagination<T>(items: T[], itemsPerPage: number = 12) {
    const currentPage = ref(1)
    
    const paginatedItems = computed(() => {
      const start = (currentPage.value - 1) * itemsPerPage
      const end = start + itemsPerPage
      return items.slice(start, end)
    })
  
    const totalItems = computed(() => items.length)
  
    const updatePage = (page: number) => {
      currentPage.value = page
      window.scrollTo({ top: 0, behavior: 'smooth' })
    }
  
    return {
      currentPage,
      paginatedItems,
      totalItems,
      itemsPerPage,
      updatePage
    }
  }