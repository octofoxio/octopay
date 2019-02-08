import * as next from 'next'
import {GraphQLServer} from "graphql-yoga";
import {gql} from "apollo-boost";
import { PaymentClient } from './generated/payment_grpc_pb'
import {Payment_ReferenceResolver, PaymentListResolver} from "./resolvers/transactionList";
const grpc = require('grpc')

const dev = process.env.NODE_ENV !== 'production'
const app = next({dev})

const paymentClient = new PaymentClient("localhost:3009", grpc.credentials.createInsecure())
const server = new GraphQLServer({
    context: (ctx) => {
        return {
            ...ctx,
            paymentClient,
        }
    },
    resolvers: {
        Payment: {
            reference: Payment_ReferenceResolver,
        },
        Query :{
            payments: PaymentListResolver,
        }

    },
    typeDefs: gql`
            scalar JSON
            type Query {
                payments(limit: Int = 10, offset: Int): [Payment!]!
            }
            type PaymentHistory {
                memo: String
                status: String!
            } 
            type PaymentProvider {
                id: String!
                type: String!
            } 
            
            type PaymentBarcodeReference {
                code: String!
                format: String!
            }
            type PaymentRedirectReference {
                url: String!
            }
            union PaymentReference = PaymentBarcodeReference | PaymentRedirectReference
            type Payment {
                id: String!
                amount: Float!
                currency: String!
                histories: [PaymentHistory!]!
                provider: PaymentProvider
                reference: PaymentReference
            }
        `
})

server.start({
    port: 3001,
    endpoint: "/graphql",
    playground: "/graphql"
}, () => {
    console.log("Server start at 3001")
})

app.prepare().then(() => {
    const handle = app.getRequestHandler()
    server.express.get("*", (req, res) => {
        return handle(req, res)
    })
})



