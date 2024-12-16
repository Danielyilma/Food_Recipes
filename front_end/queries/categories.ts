export const getCategorisQuery = gql`
    query {
        categories {
            id
            name    
            description
            image
            recipes_aggregate {
                aggregate {
                    count
                }
            }
        }
    }
`