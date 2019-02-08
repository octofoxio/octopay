// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var payment_pb = require('./payment_pb.js');

function serialize_agent_ConfirmPaymentCashInInput(arg) {
  if (!(arg instanceof payment_pb.ConfirmPaymentCashInInput)) {
    throw new Error('Expected argument of type agent.ConfirmPaymentCashInInput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_agent_ConfirmPaymentCashInInput(buffer_arg) {
  return payment_pb.ConfirmPaymentCashInInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_agent_ConfirmPaymentCashInOutput(arg) {
  if (!(arg instanceof payment_pb.ConfirmPaymentCashInOutput)) {
    throw new Error('Expected argument of type agent.ConfirmPaymentCashInOutput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_agent_ConfirmPaymentCashInOutput(buffer_arg) {
  return payment_pb.ConfirmPaymentCashInOutput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_agent_GetTransactionListInput(arg) {
  if (!(arg instanceof payment_pb.GetTransactionListInput)) {
    throw new Error('Expected argument of type agent.GetTransactionListInput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_agent_GetTransactionListInput(buffer_arg) {
  return payment_pb.GetTransactionListInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_agent_GetTransactionListOutput(arg) {
  if (!(arg instanceof payment_pb.GetTransactionListOutput)) {
    throw new Error('Expected argument of type agent.GetTransactionListOutput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_agent_GetTransactionListOutput(buffer_arg) {
  return payment_pb.GetTransactionListOutput.deserializeBinary(new Uint8Array(buffer_arg));
}


var PaymentService = exports.PaymentService = {
  confirmPaymentCashIn: {
    path: '/agent.Payment/ConfirmPaymentCashIn',
    requestStream: false,
    responseStream: false,
    requestType: payment_pb.ConfirmPaymentCashInInput,
    responseType: payment_pb.ConfirmPaymentCashInOutput,
    requestSerialize: serialize_agent_ConfirmPaymentCashInInput,
    requestDeserialize: deserialize_agent_ConfirmPaymentCashInInput,
    responseSerialize: serialize_agent_ConfirmPaymentCashInOutput,
    responseDeserialize: deserialize_agent_ConfirmPaymentCashInOutput,
  },
  getTransactionList: {
    path: '/agent.Payment/GetTransactionList',
    requestStream: false,
    responseStream: false,
    requestType: payment_pb.GetTransactionListInput,
    responseType: payment_pb.GetTransactionListOutput,
    requestSerialize: serialize_agent_GetTransactionListInput,
    requestDeserialize: deserialize_agent_GetTransactionListInput,
    responseSerialize: serialize_agent_GetTransactionListOutput,
    responseDeserialize: deserialize_agent_GetTransactionListOutput,
  },
};

exports.PaymentClient = grpc.makeGenericClientConstructor(PaymentService);
