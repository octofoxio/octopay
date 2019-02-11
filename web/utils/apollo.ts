import 'isomorphic-fetch'
import ApolloClient from 'apollo-boost'
import * as React from 'react'
import {Transaction} from "../generated/payment_pb";

export const apolloClient = new ApolloClient({
    uri: typeof window === "undefined" ? "http://localhost:3001/graphql" : "/graphql",
})

export const ApolloClientContext = React.createContext(apolloClient)



export function TransactionGRPCToGraphQL(tx?: Transaction) {
    if (!tx  || typeof tx === "undefined") {
        return null
    }
    let p = tx.getPaymentprovider()
    let provider: { type: string, id: string } | null = null
    if (typeof p !== "undefined") {
        provider = {
            id: p.getId(),
            type: p.getType(),
        }
    }
    return {
        id: tx.getId(),
        amount: tx.getAmount(),
        currency: tx.getCurrency(),
        cashInReference: tx.getCashinreference(),
        cashInType: tx.getCashintype(),
        provider,
        histories: tx.getHistoryList().map(history => {
            return {
                status: history.getStatus(),
                memo: history.getMemo(),
            }
        }),
    }

}
