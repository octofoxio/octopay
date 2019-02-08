import 'isomorphic-fetch'
import ApolloClient from 'apollo-boost'
import * as React from 'react'

export const apolloClient = new ApolloClient({
    uri: typeof window === "undefined" ? "http://localhost:3001/graphql" : "/graphql",
})

export const ApolloClientContext = React.createContext(apolloClient)
