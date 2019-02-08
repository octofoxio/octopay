import * as React from 'react'
import {ApolloProvider} from 'react-apollo'
import {apolloClient, ApolloClientContext} from "../utils/apollo";
import {PaymentCashInReference, UseTransactionList} from "../hooks/transaction-list";


const NavigationBar = () => {
    return (
        <div className={"navigation-bar"}>

        </div>
    )
}

const MenuBar = () => {
    return (
        <div className={"menu-bar"}>
            <div className={"menu-bar__list py-3"}>
                <div className={"menu-bar__item d-flex align-items-center px-3"}>
                    <div className={"icon-placeholder mr-1"}/>
                    <div className={""}>
                        {"Transactions"}
                    </div>
                </div>
            </div>

        </div>
    )
}

function getPaymentReference(paymentRef?: PaymentCashInReference | null) {
            if (!paymentRef) {
                return "-"
            }
            if (paymentRef.__typename === "PaymentBarcodeReference") {
                return `${paymentRef.code} (${paymentRef.format})`
            } else {
                return paymentRef.url
            }
}

const TransactionList = () => {
    const {txList} = UseTransactionList()
    return (
        <div className={"tx-list section py-3 px-3 my-3"}>
            <table className="table">
                <thead className={"tx-header"}>
                <tr className="">
                    <th className={"px-3"}>
                        {"ID"}
                    </th>
                    <th className={"px-3"}>
                        {"Payment provider"}
                    </th>

                    <th className={"px-3"}>
                        {"Cash-in type"}
                    </th>
                    <th className={"px-3"}>
                        {"Cash-in reference"}
                    </th>
                    <th className={"px-3"}>
                        {"Currency"}
                    </th>
                    <th className={"px-3"}>
                        {"Amount"}
                    </th>
                    <th className={"px-3"}>
                        {"Status"}
                    </th>
                </tr>
                </thead>
                <tbody>
                {
                    txList.map(tx => {
                        console.log(tx.reference)
                        return (
                            <tr key={tx.id}>
                                <td className={"px-3"}>
                                    {tx.id}
                                </td>
                                <td className={"px-3"}>
                                    {tx.provider.id}
                                </td>
                                <td>
                                    {tx.reference ? tx.reference.__typename : null}
                                </td>
                                    <td className={"px-3"}>
                                    {getPaymentReference(tx.reference)}
                                </td>
                                    <td className={"px-3"}>
                                    {tx.currency}
                                </td>
                                    <td className={"px-3"}>
                                    {tx.amount}
                                </td>
                                    <td className={"px-3"}>
                                    {tx.histories[tx.histories.length - 1].status}
                                </td>
                            </tr>
                        )

                    })
                }
                </tbody>
            </table>
        </div>
    )

}

export default () => {
    return (
        <ApolloClientContext.Provider value={apolloClient}>
            <ApolloProvider client={apolloClient}>
                <div className={"wrap d-flex flex-column"}>
                    <NavigationBar/>
                    <div className={"flex-grow-1 d-flex content"}>
                        <MenuBar/>
                        <div className={"flex-grow-1 px-2 py-2"}>
                            <TransactionList/>
                        </div>
                    </div>
                </div>
            </ApolloProvider>
        </ApolloClientContext.Provider>
    )
}