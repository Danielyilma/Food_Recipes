export const getRecipesQuery = gql`
  query ($limit: Int!, $offset: Int!) {
    recipes(limit: $limit, offset: $offset) {
      id
      title
      description
      is_paid
      prep_time
      thumbnail
      user {
        username
        profile_image
      }
      category {
        name
      }
      created_at
    }
    recipes_aggregate {
      aggregate {
        count
      }
    }
  }
`;

export const getRecipeQuery = gql`
  query GetRecipeById($id: Int!) {
    recipes_by_pk(id: $id) {
      user {
        profile_image
        username
      }
      category {
        description
        id
        name
      }
      title
      category_id
      description
      id
      created_at
      is_paid
      prep_time
      price
      recipe_images {
        image_url
      }
      steps {
        id
        step_number
        instruction
        recipe_images {
          image_url
        }
      }
      recipe_ingredients {
        unit
        quantity
        ingredient {
          name
        }
      }
    }
  }

`;

export const getFilteredRecipesQuery = gql`
  query getdata(
    $search: String
    $categories: [String]
    $prepTimeQuick: Boolean
    $prepTimeMedium: Boolean
    $prepTimeLong: Boolean
    $creator: String
  ) {
    recipes(
      where: {
        _or: [
          {
            _or: [
              { title: { _ilike: $search } }
              { description: { _ilike: $search } }
            ]
          }
          { category: { name: { _in: $categories } } }
          { user: { username: { _ilike: $creator } } }
        ]
      }
    ) {
      title
      description
      is_paid
      prep_time
      thumbnail
      user {
        username
        profile_image
      }
      category {
        name
      }
      created_at
    }
  }
`;

export const insertRecipeCommentQuery = gql`
  mutation InsertComment($recipeId: Int!, $content: String!, $userId: Int!) {
    insert_comments(objects: {recipe_id: $recipeId, content: $content, user_id: $userId}) {
      returning {
        id
      }
    }
  }
`;


export const likeRecipeQuery = gql`
  mutation LikeRecipe($recipeId: Int!, $userId: Int!) {
    insert_likes(
      objects: { recipe_id: $recipeId, user_id: $userId }
      on_conflict: { constraint: likes_pkey, update_columns: [] }
    ) {
      returning {
        id
      }
    }
  }
`;

export const dislikeRecipeQuery = gql`
  mutation DislikeRecipe($recipeId: Int!, $userId: Int!) {
  delete_likes(where: { recipe_id: { _eq: $recipeId }, user_id: { _eq: $userId } }) {
    returning {
      id
    }
  }
}
`;

export const getlikesByRecipeAndUserQuery = gql`
  query GetLikesByRecipeAndUser($recipeId: Int!, $userId: Int!) {
    likes(where: {recipe_id: {_eq: $recipeId}, user_id: {_eq: $userId}}) {
      id
    }
  }
`;

export const getCommentQuery = gql`
  query GetCommentsByRecipe($recipeId: Int!, $userId: Int!) {
    comments(where: { recipe_id: { _eq: $recipeId }, user_id: {_eq: $userId} }) {
      id
      user_id
      user {
        username
        profile_image
      }
      content
      created_at
    }
  }
`;

export const getLikeCountandislikedQuery = gql`
  query GetLikeInfo($recipeId: Int!, $userId: Int!) {
  likes_aggregate(where: {recipe_id: {_eq: $recipeId}}) {
    aggregate {
      count
    }
  }
  likes(where: {recipe_id: {_eq: $recipeId}, user_id: {_eq: $userId}}) {
    id
  }
}
`;

export const getUserRecipeQuery = gql`
  query getUserRecipe($id : Int!){
    recipes(where: {user_id: {_eq: $id}}){
      id
      title
      description
      thumbnail
      category {
        name
      }
      prep_time
      created_at
      likes_aggregate {
        aggregate{
          count
        }
      }
    }
  }
`;

export const deleteRecipeQuery = gql`
  mutation DeleteRecipe($id: Int!) {
    delete_recipes(where: {id: {_eq: $id}}) {
      returning {
        id
      }
    }
  }
`;