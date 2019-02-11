// package: agent
// file: payment.proto

import * as grpc from 'grpc';
import * as payment_pb from './payment_pb';

interface IPaymentService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  confirmPaymentCashIn: IPaymentService_IConfirmPaymentCashIn;
  getTransactionList: IPaymentService_IGetTransactionList;
  getTransaction: IPaymentService_IGetTransaction;
  updatePaymentStatus: IPaymentService_IUpdatePaymentStatus;
}

interface IPaymentService_IConfirmPaymentCashIn {
  path: string; // "/agent.Payment/ConfirmPaymentCashIn"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<payment_pb.ConfirmPaymentCashInInput>;
  requestDeserialize: grpc.deserialize<payment_pb.ConfirmPaymentCashInInput>;
  responseSerialize: grpc.serialize<payment_pb.ConfirmPaymentCashInOutput>;
  responseDeserialize: grpc.deserialize<payment_pb.ConfirmPaymentCashInOutput>;
}

interface IPaymentService_IGetTransactionList {
  path: string; // "/agent.Payment/GetTransactionList"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<payment_pb.GetTransactionListInput>;
  requestDeserialize: grpc.deserialize<payment_pb.GetTransactionListInput>;
  responseSerialize: grpc.serialize<payment_pb.GetTransactionListOutput>;
  responseDeserialize: grpc.deserialize<payment_pb.GetTransactionListOutput>;
}

interface IPaymentService_IGetTransaction {
  path: string; // "/agent.Payment/GetTransaction"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<payment_pb.GetTransactionInput>;
  requestDeserialize: grpc.deserialize<payment_pb.GetTransactionInput>;
  responseSerialize: grpc.serialize<payment_pb.GetTransactionOutput>;
  responseDeserialize: grpc.deserialize<payment_pb.GetTransactionOutput>;
}

interface IPaymentService_IUpdatePaymentStatus {
  path: string; // "/agent.Payment/UpdatePaymentStatus"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<payment_pb.UpdatePaymentStatusInput>;
  requestDeserialize: grpc.deserialize<payment_pb.UpdatePaymentStatusInput>;
  responseSerialize: grpc.serialize<payment_pb.UpdatePaymentStatusOutput>;
  responseDeserialize: grpc.deserialize<payment_pb.UpdatePaymentStatusOutput>;
}

export const PaymentService: IPaymentService;
export interface IPaymentServer {
  confirmPaymentCashIn: grpc.handleUnaryCall<payment_pb.ConfirmPaymentCashInInput, payment_pb.ConfirmPaymentCashInOutput>;
  getTransactionList: grpc.handleUnaryCall<payment_pb.GetTransactionListInput, payment_pb.GetTransactionListOutput>;
  getTransaction: grpc.handleUnaryCall<payment_pb.GetTransactionInput, payment_pb.GetTransactionOutput>;
  updatePaymentStatus: grpc.handleUnaryCall<payment_pb.UpdatePaymentStatusInput, payment_pb.UpdatePaymentStatusOutput>;
}

export interface IPaymentClient {
  confirmPaymentCashIn(request: payment_pb.ConfirmPaymentCashInInput, callback: (error: Error | null, response: payment_pb.ConfirmPaymentCashInOutput) => void): grpc.ClientUnaryCall;
  confirmPaymentCashIn(request: payment_pb.ConfirmPaymentCashInInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.ConfirmPaymentCashInOutput) => void): grpc.ClientUnaryCall;
  getTransactionList(request: payment_pb.GetTransactionListInput, callback: (error: Error | null, response: payment_pb.GetTransactionListOutput) => void): grpc.ClientUnaryCall;
  getTransactionList(request: payment_pb.GetTransactionListInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.GetTransactionListOutput) => void): grpc.ClientUnaryCall;
  getTransaction(request: payment_pb.GetTransactionInput, callback: (error: Error | null, response: payment_pb.GetTransactionOutput) => void): grpc.ClientUnaryCall;
  getTransaction(request: payment_pb.GetTransactionInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.GetTransactionOutput) => void): grpc.ClientUnaryCall;
  updatePaymentStatus(request: payment_pb.UpdatePaymentStatusInput, callback: (error: Error | null, response: payment_pb.UpdatePaymentStatusOutput) => void): grpc.ClientUnaryCall;
  updatePaymentStatus(request: payment_pb.UpdatePaymentStatusInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.UpdatePaymentStatusOutput) => void): grpc.ClientUnaryCall;
}

export class PaymentClient extends grpc.Client implements IPaymentClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  public confirmPaymentCashIn(request: payment_pb.ConfirmPaymentCashInInput, callback: (error: Error | null, response: payment_pb.ConfirmPaymentCashInOutput) => void): grpc.ClientUnaryCall;
  public confirmPaymentCashIn(request: payment_pb.ConfirmPaymentCashInInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.ConfirmPaymentCashInOutput) => void): grpc.ClientUnaryCall;
  public getTransactionList(request: payment_pb.GetTransactionListInput, callback: (error: Error | null, response: payment_pb.GetTransactionListOutput) => void): grpc.ClientUnaryCall;
  public getTransactionList(request: payment_pb.GetTransactionListInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.GetTransactionListOutput) => void): grpc.ClientUnaryCall;
  public getTransaction(request: payment_pb.GetTransactionInput, callback: (error: Error | null, response: payment_pb.GetTransactionOutput) => void): grpc.ClientUnaryCall;
  public getTransaction(request: payment_pb.GetTransactionInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.GetTransactionOutput) => void): grpc.ClientUnaryCall;
  public updatePaymentStatus(request: payment_pb.UpdatePaymentStatusInput, callback: (error: Error | null, response: payment_pb.UpdatePaymentStatusOutput) => void): grpc.ClientUnaryCall;
  public updatePaymentStatus(request: payment_pb.UpdatePaymentStatusInput, metadata: grpc.Metadata, callback: (error: Error | null, response: payment_pb.UpdatePaymentStatusOutput) => void): grpc.ClientUnaryCall;
}

