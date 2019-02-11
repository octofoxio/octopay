import {PaymentClient} from "../generated/payment_grpc_pb";
import {PaymentStatus, UpdatePaymentStatusInput} from "../generated/payment_pb";
import {TransactionGRPCToGraphQL} from "../utils/apollo";

interface UpdatePaymentStatusContext {
    paymentClient: PaymentClient
}
interface UpdatePaymentStatusArgs {
    id: string
    status: string
}

export function UpdatePaymentStatusResolver(source, args: UpdatePaymentStatusArgs, context: UpdatePaymentStatusContext) {
    return new Promise<any>((resolve, reject) => {
        const request = new UpdatePaymentStatusInput()
        request.setId(args.id)
        request.setStatus(PaymentStatus[args.status])
        request.setMemo(`by Admin@${new Date().toISOString()}`)
        context.paymentClient.updatePaymentStatus(request, function(err, resp) {
            resolve(TransactionGRPCToGraphQL(resp.getResult()))
        })

    })

}