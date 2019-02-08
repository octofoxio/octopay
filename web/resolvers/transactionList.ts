import {PaymentClient} from "../generated/payment_grpc_pb";
import {GetTransactionListInput} from "../generated/payment_pb";

interface TransactionListResolverContext {
    paymentClient: PaymentClient
}

interface TransactionListResolverArgs {
    offset: number
    limit: number
}

export function Payment_ReferenceResolver(source, args, ctx: TransactionListResolverContext) {
    try {
        if (!source.cashInReference) {
            return null
        }
        const data = JSON.parse(source.cashInReference)
        return {
            ...data,
            __typename: source.cashInType === "barcode" ? "PaymentBarcodeReference" : "",
        }
    } catch (e) {
        throw new Error("CashInReference parse error")
    }
}

export function PaymentListResolver(_, args: TransactionListResolverArgs, context: TransactionListResolverContext) {
    return new Promise((resolve, reject) => {
        const request = new GetTransactionListInput()
        request.setLimit(args.limit)
        request.setOffset(args.offset)
        context.paymentClient.getTransactionList(request, function (err, response) {
            if (err) {
                reject(err)
                return
            }
            const txList = response.getTransactionsList().map(transactionItem => {
                let p = transactionItem.getPaymentprovider()
                let provider: { type: string, id: string } | null = null
                if (typeof p !== "undefined") {
                    provider = {
                        id: p.getId(),
                        type: p.getType(),
                    }
                }

                return {
                    id: transactionItem.getId(),
                    amount: transactionItem.getAmount(),
                    currency: transactionItem.getCurrency(),
                    cashInReference: transactionItem.getCashinreference(),
                    cashInType: transactionItem.getCashintype(),
                    provider,
                    histories: transactionItem.getHistoryList().map(history => {
                        return {
                            status: history.getStatus(),
                            memo: history.getMemo(),
                        }
                    }),
                }
            })
            resolve(txList)
        })
    })
}