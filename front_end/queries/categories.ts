export const getCategorisQuery = gql`
    query {
        categories {
            id
            name    
            description
            recipes_aggregate {
                aggregate {
                    count
                }
            }
        }
    }
`