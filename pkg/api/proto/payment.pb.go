// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/proto/payment.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PaymentAgent struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentAgent) Reset()         { *m = PaymentAgent{} }
func (m *PaymentAgent) String() string { return proto.CompactTextString(m) }
func (*PaymentAgent) ProtoMessage()    {}
func (*PaymentAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{0}
}

func (m *PaymentAgent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentAgent.Unmarshal(m, b)
}
func (m *PaymentAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentAgent.Marshal(b, m, deterministic)
}
func (m *PaymentAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentAgent.Merge(m, src)
}
func (m *PaymentAgent) XXX_Size() int {
	return xxx_messageInfo_PaymentAgent.Size(m)
}
func (m *PaymentAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentAgent.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentAgent proto.InternalMessageInfo

func (m *PaymentAgent) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PaymentAgent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type History struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Memo                 string   `protobuf:"bytes,2,opt,name=Memo,proto3" json:"Memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *History) Reset()         { *m = History{} }
func (m *History) String() string { return proto.CompactTextString(m) }
func (*History) ProtoMessage()    {}
func (*History) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{1}
}

func (m *History) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_History.Unmarshal(m, b)
}
func (m *History) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_History.Marshal(b, m, deterministic)
}
func (m *History) XXX_Merge(src proto.Message) {
	xxx_messageInfo_History.Merge(m, src)
}
func (m *History) XXX_Size() int {
	return xxx_messageInfo_History.Size(m)
}
func (m *History) XXX_DiscardUnknown() {
	xxx_messageInfo_History.DiscardUnknown(m)
}

var xxx_messageInfo_History proto.InternalMessageInfo

func (m *History) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *History) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

type Transaction struct {
	ID                   string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PaymentProvider      *PaymentAgent `protobuf:"bytes,2,opt,name=PaymentProvider,proto3" json:"PaymentProvider,omitempty"`
	CashInReference      string        `protobuf:"bytes,3,opt,name=CashInReference,proto3" json:"CashInReference,omitempty"`
	Currency             string        `protobuf:"bytes,4,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Amount               float32       `protobuf:"fixed32,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	History              []*History    `protobuf:"bytes,6,rep,name=History,proto3" json:"History,omitempty"`
	CashInType           string        `protobuf:"bytes,7,opt,name=CashInType,proto3" json:"CashInType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{2}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Transaction) GetPaymentProvider() *PaymentAgent {
	if m != nil {
		return m.PaymentProvider
	}
	return nil
}

func (m *Transaction) GetCashInReference() string {
	if m != nil {
		return m.CashInReference
	}
	return ""
}

func (m *Transaction) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Transaction) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetHistory() []*History {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *Transaction) GetCashInType() string {
	if m != nil {
		return m.CashInType
	}
	return ""
}

type GetTransactionListInput struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionListInput) Reset()         { *m = GetTransactionListInput{} }
func (m *GetTransactionListInput) String() string { return proto.CompactTextString(m) }
func (*GetTransactionListInput) ProtoMessage()    {}
func (*GetTransactionListInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{3}
}

func (m *GetTransactionListInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionListInput.Unmarshal(m, b)
}
func (m *GetTransactionListInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionListInput.Marshal(b, m, deterministic)
}
func (m *GetTransactionListInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionListInput.Merge(m, src)
}
func (m *GetTransactionListInput) XXX_Size() int {
	return xxx_messageInfo_GetTransactionListInput.Size(m)
}
func (m *GetTransactionListInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionListInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionListInput proto.InternalMessageInfo

func (m *GetTransactionListInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetTransactionListInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type GetTransactionListOutput struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTransactionListOutput) Reset()         { *m = GetTransactionListOutput{} }
func (m *GetTransactionListOutput) String() string { return proto.CompactTextString(m) }
func (*GetTransactionListOutput) ProtoMessage()    {}
func (*GetTransactionListOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{4}
}

func (m *GetTransactionListOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionListOutput.Unmarshal(m, b)
}
func (m *GetTransactionListOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionListOutput.Marshal(b, m, deterministic)
}
func (m *GetTransactionListOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionListOutput.Merge(m, src)
}
func (m *GetTransactionListOutput) XXX_Size() int {
	return xxx_messageInfo_GetTransactionListOutput.Size(m)
}
func (m *GetTransactionListOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionListOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionListOutput proto.InternalMessageInfo

func (m *GetTransactionListOutput) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type ConfirmPaymentCashInInput struct {
	PaymentID            string   `protobuf:"bytes,1,opt,name=PaymentID,proto3" json:"PaymentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmPaymentCashInInput) Reset()         { *m = ConfirmPaymentCashInInput{} }
func (m *ConfirmPaymentCashInInput) String() string { return proto.CompactTextString(m) }
func (*ConfirmPaymentCashInInput) ProtoMessage()    {}
func (*ConfirmPaymentCashInInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{5}
}

func (m *ConfirmPaymentCashInInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmPaymentCashInInput.Unmarshal(m, b)
}
func (m *ConfirmPaymentCashInInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmPaymentCashInInput.Marshal(b, m, deterministic)
}
func (m *ConfirmPaymentCashInInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmPaymentCashInInput.Merge(m, src)
}
func (m *ConfirmPaymentCashInInput) XXX_Size() int {
	return xxx_messageInfo_ConfirmPaymentCashInInput.Size(m)
}
func (m *ConfirmPaymentCashInInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmPaymentCashInInput.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmPaymentCashInInput proto.InternalMessageInfo

func (m *ConfirmPaymentCashInInput) GetPaymentID() string {
	if m != nil {
		return m.PaymentID
	}
	return ""
}

type ConfirmPaymentCashInOutput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmPaymentCashInOutput) Reset()         { *m = ConfirmPaymentCashInOutput{} }
func (m *ConfirmPaymentCashInOutput) String() string { return proto.CompactTextString(m) }
func (*ConfirmPaymentCashInOutput) ProtoMessage()    {}
func (*ConfirmPaymentCashInOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_428c3b718c6302e1, []int{6}
}

func (m *ConfirmPaymentCashInOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmPaymentCashInOutput.Unmarshal(m, b)
}
func (m *ConfirmPaymentCashInOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmPaymentCashInOutput.Marshal(b, m, deterministic)
}
func (m *ConfirmPaymentCashInOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmPaymentCashInOutput.Merge(m, src)
}
func (m *ConfirmPaymentCashInOutput) XXX_Size() int {
	return xxx_messageInfo_ConfirmPaymentCashInOutput.Size(m)
}
func (m *ConfirmPaymentCashInOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmPaymentCashInOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmPaymentCashInOutput proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PaymentAgent)(nil), "agent.PaymentAgent")
	proto.RegisterType((*History)(nil), "agent.History")
	proto.RegisterType((*Transaction)(nil), "agent.Transaction")
	proto.RegisterType((*GetTransactionListInput)(nil), "agent.GetTransactionListInput")
	proto.RegisterType((*GetTransactionListOutput)(nil), "agent.GetTransactionListOutput")
	proto.RegisterType((*ConfirmPaymentCashInInput)(nil), "agent.ConfirmPaymentCashInInput")
	proto.RegisterType((*ConfirmPaymentCashInOutput)(nil), "agent.ConfirmPaymentCashInOutput")
}

func init() { proto.RegisterFile("pkg/api/proto/payment.proto", fileDescriptor_428c3b718c6302e1) }

var fileDescriptor_428c3b718c6302e1 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xd2, 0x66, 0x97, 0xce, 0x56, 0xad, 0x64, 0x0a, 0x0d, 0xa1, 0x2a, 0x21, 0xa7, 0x9c,
	0x36, 0x52, 0x10, 0x48, 0x1c, 0x38, 0x94, 0xad, 0x54, 0x22, 0x81, 0x58, 0x99, 0xe5, 0x02, 0xa7,
	0xb0, 0x38, 0x8b, 0x05, 0xb1, 0x23, 0x67, 0x82, 0x94, 0x3b, 0x7f, 0x8b, 0xff, 0x86, 0xec, 0x98,
	0x6c, 0xf6, 0xab, 0xa7, 0xcc, 0x7b, 0x9e, 0x37, 0x7e, 0xf3, 0xac, 0xc0, 0xd3, 0xea, 0xe7, 0x2a,
	0xc9, 0x2b, 0x9e, 0x54, 0x4a, 0xa2, 0x4c, 0xaa, 0xbc, 0x2d, 0x99, 0xc0, 0xa9, 0x41, 0xc4, 0xcb,
	0x57, 0x4c, 0x60, 0x94, 0xc2, 0xe9, 0xbc, 0xe3, 0x6f, 0x34, 0x26, 0x67, 0xe0, 0x66, 0xb7, 0xbe,
	0x13, 0x3a, 0xf1, 0x09, 0x75, 0xb3, 0x5b, 0x42, 0xe0, 0x78, 0xd1, 0x56, 0xcc, 0x77, 0x0d, 0x63,
	0xea, 0xe8, 0x25, 0x8c, 0xdf, 0xf1, 0x1a, 0xa5, 0x6a, 0xc9, 0x63, 0x18, 0x7d, 0xc2, 0x1c, 0x9b,
	0xda, 0x4a, 0x2c, 0xd2, 0xb2, 0x0f, 0xac, 0x94, 0xff, 0x65, 0xba, 0x8e, 0xfe, 0xb8, 0x30, 0x59,
	0xa8, 0x5c, 0xd4, 0xf9, 0x12, 0xb9, 0x14, 0x3b, 0x57, 0xbd, 0x81, 0x73, 0x6b, 0x65, 0xae, 0xe4,
	0x6f, 0xfe, 0x9d, 0x29, 0x23, 0x9f, 0xa4, 0x0f, 0xa7, 0xc6, 0xeb, 0x74, 0x68, 0x94, 0x6e, 0xf7,
	0x92, 0x18, 0xce, 0x67, 0x79, 0xfd, 0x23, 0x13, 0x94, 0x15, 0x4c, 0x31, 0xb1, 0x64, 0xfe, 0x91,
	0x99, 0xbd, 0x4d, 0x93, 0x00, 0x1e, 0xcc, 0x1a, 0xa5, 0xeb, 0xd6, 0x3f, 0x36, 0x2d, 0x3d, 0xd6,
	0x0b, 0xdd, 0x94, 0xb2, 0x11, 0xe8, 0x7b, 0xa1, 0x13, 0xbb, 0xd4, 0x22, 0x12, 0xf7, 0x3b, 0xfb,
	0xa3, 0xf0, 0x28, 0x9e, 0xa4, 0x67, 0xd6, 0x94, 0x65, 0x69, 0x1f, 0xc9, 0x35, 0x40, 0x77, 0xa1,
	0xc9, 0x6d, 0x6c, 0xe6, 0x0f, 0x98, 0xe8, 0x0e, 0x2e, 0xef, 0x18, 0x0e, 0x82, 0x78, 0xcf, 0x6b,
	0xcc, 0x44, 0xd5, 0x20, 0xb9, 0x00, 0xef, 0x17, 0x2f, 0x39, 0x9a, 0x50, 0x3c, 0xda, 0x01, 0x6d,
	0x49, 0x16, 0x45, 0xcd, 0xd0, 0xc4, 0xe1, 0x51, 0x8b, 0x22, 0x0a, 0xfe, 0xee, 0xa0, 0x8f, 0x0d,
	0xea, 0x49, 0xaf, 0xe0, 0x14, 0xd7, 0x07, 0xfa, 0x75, 0xb4, 0x67, 0x62, 0x3d, 0x0f, 0x34, 0x74,
	0xa3, 0x2f, 0x7a, 0x0d, 0x4f, 0x66, 0x52, 0x14, 0x5c, 0x95, 0x36, 0xde, 0xce, 0x78, 0x67, 0xef,
	0x0a, 0x4e, 0x2c, 0xdb, 0xbf, 0xdb, 0x9a, 0x88, 0xae, 0x20, 0xd8, 0x27, 0xed, 0x0c, 0xa5, 0x7f,
	0x1d, 0x18, 0x5b, 0x9e, 0x7c, 0x85, 0x8b, 0x7d, 0x9d, 0x24, 0xb4, 0xf6, 0x0e, 0x3a, 0x08, 0x9e,
	0xdf, 0xd3, 0x61, 0x37, 0xff, 0x0c, 0x64, 0x37, 0x15, 0x72, 0x6d, 0x85, 0x07, 0x92, 0x0f, 0x9e,
	0x1d, 0x3c, 0xef, 0xc6, 0xbe, 0xbd, 0xfc, 0xf2, 0xa8, 0x96, 0x05, 0x0a, 0x86, 0xc9, 0xc6, 0x5f,
	0xf5, 0x6d, 0x64, 0x3e, 0x2f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xba, 0x02, 0xf8, 0x90, 0x6d,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentClient interface {
	ConfirmPaymentCashIn(ctx context.Context, in *ConfirmPaymentCashInInput, opts ...grpc.CallOption) (*ConfirmPaymentCashInOutput, error)
	GetTransactionList(ctx context.Context, in *GetTransactionListInput, opts ...grpc.CallOption) (*GetTransactionListOutput, error)
}

type paymentClient struct {
	cc *grpc.ClientConn
}

func NewPaymentClient(cc *grpc.ClientConn) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) ConfirmPaymentCashIn(ctx context.Context, in *ConfirmPaymentCashInInput, opts ...grpc.CallOption) (*ConfirmPaymentCashInOutput, error) {
	out := new(ConfirmPaymentCashInOutput)
	err := c.cc.Invoke(ctx, "/agent.Payment/ConfirmPaymentCashIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetTransactionList(ctx context.Context, in *GetTransactionListInput, opts ...grpc.CallOption) (*GetTransactionListOutput, error) {
	out := new(GetTransactionListOutput)
	err := c.cc.Invoke(ctx, "/agent.Payment/GetTransactionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
type PaymentServer interface {
	ConfirmPaymentCashIn(context.Context, *ConfirmPaymentCashInInput) (*ConfirmPaymentCashInOutput, error)
	GetTransactionList(context.Context, *GetTransactionListInput) (*GetTransactionListOutput, error)
}

func RegisterPaymentServer(s *grpc.Server, srv PaymentServer) {
	s.RegisterService(&_Payment_serviceDesc, srv)
}

func _Payment_ConfirmPaymentCashIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmPaymentCashInInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).ConfirmPaymentCashIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Payment/ConfirmPaymentCashIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).ConfirmPaymentCashIn(ctx, req.(*ConfirmPaymentCashInInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetTransactionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionListInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetTransactionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Payment/GetTransactionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetTransactionList(ctx, req.(*GetTransactionListInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfirmPaymentCashIn",
			Handler:    _Payment_ConfirmPaymentCashIn_Handler,
		},
		{
			MethodName: "GetTransactionList",
			Handler:    _Payment_GetTransactionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/proto/payment.proto",
}
