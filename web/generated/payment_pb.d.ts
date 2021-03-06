// package: agent
// file: payment.proto

import * as jspb from 'google-protobuf';

export class PaymentAgent extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getType(): string;
  setType(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PaymentAgent.AsObject;
  static toObject(includeInstance: boolean, msg: PaymentAgent): PaymentAgent.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PaymentAgent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PaymentAgent;
  static deserializeBinaryFromReader(message: PaymentAgent, reader: jspb.BinaryReader): PaymentAgent;
}

export namespace PaymentAgent {
  export type AsObject = {
    id: string,
    type: string,
  }
}

export class History extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): void;

  getMemo(): string;
  setMemo(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): History.AsObject;
  static toObject(includeInstance: boolean, msg: History): History.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: History, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): History;
  static deserializeBinaryFromReader(message: History, reader: jspb.BinaryReader): History;
}

export namespace History {
  export type AsObject = {
    status: string,
    memo: string,
  }
}

export class Transaction extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasPaymentprovider(): boolean;
  clearPaymentprovider(): void;
  getPaymentprovider(): PaymentAgent | undefined;
  setPaymentprovider(value?: PaymentAgent): void;

  getCashinreference(): string;
  setCashinreference(value: string): void;

  getCurrency(): string;
  setCurrency(value: string): void;

  getAmount(): number;
  setAmount(value: number): void;

  clearHistoryList(): void;
  getHistoryList(): Array<History>;
  setHistoryList(value: Array<History>): void;
  addHistory(value?: History, index?: number): History;

  getCashintype(): string;
  setCashintype(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Transaction.AsObject;
  static toObject(includeInstance: boolean, msg: Transaction): Transaction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Transaction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Transaction;
  static deserializeBinaryFromReader(message: Transaction, reader: jspb.BinaryReader): Transaction;
}

export namespace Transaction {
  export type AsObject = {
    id: string,
    paymentprovider?: PaymentAgent.AsObject,
    cashinreference: string,
    currency: string,
    amount: number,
    historyList: Array<History.AsObject>,
    cashintype: string,
  }
}

export class UpdatePaymentStatusInput extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getStatus(): PaymentStatus;
  setStatus(value: PaymentStatus): void;

  getMemo(): string;
  setMemo(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePaymentStatusInput.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePaymentStatusInput): UpdatePaymentStatusInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePaymentStatusInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePaymentStatusInput;
  static deserializeBinaryFromReader(message: UpdatePaymentStatusInput, reader: jspb.BinaryReader): UpdatePaymentStatusInput;
}

export namespace UpdatePaymentStatusInput {
  export type AsObject = {
    id: string,
    status: PaymentStatus,
    memo: string,
  }
}

export class UpdatePaymentStatusOutput extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Transaction | undefined;
  setResult(value?: Transaction): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePaymentStatusOutput.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePaymentStatusOutput): UpdatePaymentStatusOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePaymentStatusOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePaymentStatusOutput;
  static deserializeBinaryFromReader(message: UpdatePaymentStatusOutput, reader: jspb.BinaryReader): UpdatePaymentStatusOutput;
}

export namespace UpdatePaymentStatusOutput {
  export type AsObject = {
    result?: Transaction.AsObject,
  }
}

export class GetTransactionInput extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionInput.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionInput): GetTransactionInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetTransactionInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionInput;
  static deserializeBinaryFromReader(message: GetTransactionInput, reader: jspb.BinaryReader): GetTransactionInput;
}

export namespace GetTransactionInput {
  export type AsObject = {
    id: string,
  }
}

export class GetTransactionOutput extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Transaction | undefined;
  setResult(value?: Transaction): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionOutput.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionOutput): GetTransactionOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetTransactionOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionOutput;
  static deserializeBinaryFromReader(message: GetTransactionOutput, reader: jspb.BinaryReader): GetTransactionOutput;
}

export namespace GetTransactionOutput {
  export type AsObject = {
    result?: Transaction.AsObject,
  }
}

export class GetTransactionListInput extends jspb.Message {
  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionListInput.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionListInput): GetTransactionListInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetTransactionListInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionListInput;
  static deserializeBinaryFromReader(message: GetTransactionListInput, reader: jspb.BinaryReader): GetTransactionListInput;
}

export namespace GetTransactionListInput {
  export type AsObject = {
    limit: number,
    offset: number,
  }
}

export class GetTransactionListOutput extends jspb.Message {
  clearTransactionsList(): void;
  getTransactionsList(): Array<Transaction>;
  setTransactionsList(value: Array<Transaction>): void;
  addTransactions(value?: Transaction, index?: number): Transaction;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionListOutput.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionListOutput): GetTransactionListOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetTransactionListOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionListOutput;
  static deserializeBinaryFromReader(message: GetTransactionListOutput, reader: jspb.BinaryReader): GetTransactionListOutput;
}

export namespace GetTransactionListOutput {
  export type AsObject = {
    transactionsList: Array<Transaction.AsObject>,
  }
}

export class ConfirmPaymentCashInInput extends jspb.Message {
  getPaymentid(): string;
  setPaymentid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConfirmPaymentCashInInput.AsObject;
  static toObject(includeInstance: boolean, msg: ConfirmPaymentCashInInput): ConfirmPaymentCashInInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ConfirmPaymentCashInInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConfirmPaymentCashInInput;
  static deserializeBinaryFromReader(message: ConfirmPaymentCashInInput, reader: jspb.BinaryReader): ConfirmPaymentCashInInput;
}

export namespace ConfirmPaymentCashInInput {
  export type AsObject = {
    paymentid: string,
  }
}

export class ConfirmPaymentCashInOutput extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConfirmPaymentCashInOutput.AsObject;
  static toObject(includeInstance: boolean, msg: ConfirmPaymentCashInOutput): ConfirmPaymentCashInOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ConfirmPaymentCashInOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConfirmPaymentCashInOutput;
  static deserializeBinaryFromReader(message: ConfirmPaymentCashInOutput, reader: jspb.BinaryReader): ConfirmPaymentCashInOutput;
}

export namespace ConfirmPaymentCashInOutput {
  export type AsObject = {
  }
}

export enum PaymentStatus {
  INITIAL = 0,
  READY_TO_CASH_IN = 1,
  CASH_IN_CONFIRM = 2,
  CALLBACK_ATTEMPT_SUCCESS = 3,
  CALLBACK_ATTEMPT_FAILED = 4,
  FAILED = 5,
}

