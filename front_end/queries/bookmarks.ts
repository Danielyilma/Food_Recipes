

export const addBookmarksQuery = gql`
    mutation InsertBookmark($recipeId: Int!, $userId: Int!) {
        insert_bookmarks(objects: {recipe_id: $recipeId, user_id: $userId}) {
            affected_rows
        }
    }
`;

export const getBookmarkedrecipes = gql`
  query GetBookmarks($userId: Int!) {
    bookmarks(where: {user_id: {_eq: $userId}}) {
      recipe {
        category {
          name
          description
        }
        id
        description
        is_paid
        title
        thumbnail
        prep_time
        user {
          profile_image
          username
        }
      }
    }
  }
`;