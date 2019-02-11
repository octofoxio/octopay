import {PaymentCashInReference, UseTransactionList} from "../hooks/transaction-list";
import * as React from "react";
import {majorScale, Pane, Table} from 'evergreen-ui'
import {  } from "./PaymentDialog";

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

export const TransactionList = () => {
    const {txList} = UseTransactionList()
    return (
        <>
        <Pane
            border
            backgroundColor="var(--white)"
            padding={majorScale(2)}
            margin={majorScale(3)}
            display={"inline-block"}
        >
            <Table className="table">
                <Table.Head className={"tx-header"}>
                    <Table.SearchHeaderCell
                        placeholder="ID"
                    />
                    <Table.SearchHeaderCell
                        placeholder={"Payment provider"}
                    />
                    <Table.SearchHeaderCell
                        placeholder={"Cash-in type"}
                    />
                    <Table.SearchHeaderCell
                        placeholder={"Cash-in ref"}
                    />
                    <Table.SearchHeaderCell
                        placeholder={"Currency"}
                    />
                    <Table.SearchHeaderCell
                        placeholder={"Amount"}
                    />
                    <Table.SearchHeaderCell
                        placeholder={"status"}/>
                </Table.Head>
                <Table.Body>
                    {
                        txList.map(tx => {
                            return (
                                <Table.Row isSelectable key={tx.id}>
                                    <Table.TextCell className={"px-3"}>
                                        {tx.id}
                                    </Table.TextCell>
                                    <Table.TextCell className={"px-3"}>
                                        {tx.provider.id}
                                    </Table.TextCell>
                                    <Table.TextCell>
                                        {tx.reference ? tx.reference.__typename : null}
                                    </Table.TextCell>
                                    <Table.TextCell className={"px-3"}>
                                        {getPaymentReference(tx.reference)}
                                    </Table.TextCell>
                                    <Table.TextCell className={"px-3"}>
                                        {tx.currency}
                                    </Table.TextCell>
                                    <Table.TextCell className={"px-3"}>
                                        {tx.amount}
                                    </Table.TextCell>
                                    <Table.TextCell className={"px-3"}>
                                        {tx.histories[tx.histories.length - 1].status}
                                    </Table.TextCell>
                                </Table.Row>
                            )

                        })
                    }
                </Table.Body>
            </Table>
        </Pane>
            </>
    )

}