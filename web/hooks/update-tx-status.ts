import * as React from 'react'
import {ApolloClientContext} from "../utils/apollo";
import {gql} from "apollo-boost";

const UpdatePaymentStatusMutation = gql`
    mutation ($ID: String!, $Status: String!) {
        updatePaymentStatus(id: $ID, status: $Status) {
            id 
        }
    }
`

export function UseUpdatePaymentStatus() {
    const [isLoading, setIsLoading] = React.useState(false)
    const client = React.useContext(ApolloClientContext)
    return {
        mutate: () => {
            setIsLoading(true)
            client.mutate({
                mutation:UpdatePaymentStatusMutation,
            })
        },
        isLoading
    }
}