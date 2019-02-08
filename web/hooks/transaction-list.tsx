import * as React from 'react'
import {ApolloClientContext} from "../utils/apollo";
import {gql} from "apollo-boost";
interface PaymentCashInReferenceBarcode {
    __typename: "PaymentBarcodeReference"
    code: string
    format: string
}
interface PaymentCashInReferenceRedirect {
    __typename: "PaymentRedirectReference"
    url: string
}
export type PaymentCashInReference = PaymentCashInReferenceBarcode | PaymentCashInReferenceRedirect

export interface PaymentListResultItem {
    id: string
    amount: number
    currency: string
    histories: {memo:string,status:string}[]
    provider: {
        id: string
        type: string
    }
    reference: PaymentCashInReference | null
}
const PaymentListQuery = gql`
query {
  result: payments {
    id
    amount
    currency
    histories{
      memo
      status
    }
    provider {
      id
      type
    }
   	reference {
      __typename 
      ... on PaymentBarcodeReference {
        code
        format
      }
    }
  }
}
`

export function UseTransactionList() {
    const [txList, setTxList] = React.useState<PaymentListResultItem[]>([])
    const [loading, setLoading] = React.useState(false)
    const client = React.useContext(ApolloClientContext)

    React.useEffect(() => {

       client.watchQuery<any>({
           query:PaymentListQuery,
       }).subscribe(({data, loading}) => {
            setLoading(loading)
            if (!loading) {
                setTxList(data.result)
            }
       })
    },[])

    return {
        txList,
        loading,
    }
}