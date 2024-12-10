export const loginQuery = gql`
  mutation ($email: String!, $password: String!) {
    login(arg: { email: $email, password: $password }) {
      token
    }
  }
`;

export const signupQuery = gql`
  mutation (
    $email: String!
    $username: String!
    $phone_number: String!
    $password: String!
  ) {
    signup(
      arg1: {
        email: $email
        username: $username
        phone_number: $phone_number
        password: $password
      }
    ) {
      token
    }
  }
`;

export const getUserQuery = gql`
  query ($id: Int!) {
    users(where: { id: { _eq: $id } }) {
      id
      username
      email
      bio
      tag_name
      recipes_aggregate {
        aggregate {
          count
        }
      }
      profile_image
      phone_number
      created_at
    }
  }
`;

export const getUsersQuery = gql`
  query ($username: String ) {
    users (username: $username ){
      username
      email
      profile_image
      phone_number
    }
  }
`;

export const updateUserQuery = gql`
  mutation UpdateUser($id: Int!, $bio: String!, $profileImage: String!) {
    update_users_by_pk(pk_columns: {id: $id}, _set: {bio: $bio, profile_image: $profileImage}) {
      id
    }
  }
`;