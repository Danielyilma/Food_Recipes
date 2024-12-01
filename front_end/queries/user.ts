export const loginQuery = gql`
  mutation ($email: String!, $password: String!) {
    login(arg: { email: $email, password: $password }) {
      token
    }
  }
`;

export const getUserQuery = gql`
  query ($id: Int!) {
    users(where: { id: { _eq: $id } }) {
      username
      email
      profile_image
      phone_number
    }
  }
`;
