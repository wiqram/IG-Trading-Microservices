// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protos/uibackend/uibackend.proto

package uibackend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProductsRequest) Reset() {
	*x = ProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsRequest) ProtoMessage() {}

func (x *ProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsRequest.ProtoReflect.Descriptor instead.
func (*ProductsRequest) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{0}
}

type ProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductsResponse) Reset() {
	*x = ProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResponse) ProtoMessage() {}

func (x *ProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResponse.ProtoReflect.Descriptor instead.
func (*ProductsResponse) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{1}
}

func (x *ProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId        string `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	ProductName      string `protobuf:"bytes,3,opt,name=productName,proto3" json:"productName,omitempty"`
	SortCode         string `protobuf:"bytes,4,opt,name=sortCode,proto3" json:"sortCode,omitempty"`
	AccountNumber    string `protobuf:"bytes,5,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	CurrentBalance   int32  `protobuf:"varint,6,opt,name=currentBalance,proto3" json:"currentBalance,omitempty"`
	AvailableBalance int32  `protobuf:"varint,7,opt,name=availableBalance,proto3" json:"availableBalance,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Product) GetSortCode() string {
	if x != nil {
		return x.SortCode
	}
	return ""
}

func (x *Product) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Product) GetCurrentBalance() int32 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

func (x *Product) GetAvailableBalance() int32 {
	if x != nil {
		return x.AvailableBalance
	}
	return 0
}

type TransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	FilterBy  string `protobuf:"bytes,2,opt,name=filterBy,proto3" json:"filterBy,omitempty"`
}

func (x *TransactionsRequest) Reset() {
	*x = TransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsRequest) ProtoMessage() {}

func (x *TransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsRequest.ProtoReflect.Descriptor instead.
func (*TransactionsRequest) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionsRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *TransactionsRequest) GetFilterBy() string {
	if x != nil {
		return x.FilterBy
	}
	return ""
}

type TransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionsResponse) Reset() {
	*x = TransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsResponse) ProtoMessage() {}

func (x *TransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsResponse.ProtoReflect.Descriptor instead.
func (*TransactionsResponse) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId        string             `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	DateOn           string             `protobuf:"bytes,3,opt,name=dateOn,proto3" json:"dateOn,omitempty"`
	DebitCredit      string             `protobuf:"bytes,4,opt,name=debitCredit,proto3" json:"debitCredit,omitempty"`
	TransDescription string             `protobuf:"bytes,5,opt,name=transDescription,proto3" json:"transDescription,omitempty"`
	Details          string             `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
	Amount           int32              `protobuf:"varint,7,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionList  []*TransactionList `protobuf:"bytes,8,rep,name=transactionList,proto3" json:"transactionList,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Transaction) GetDateOn() string {
	if x != nil {
		return x.DateOn
	}
	return ""
}

func (x *Transaction) GetDebitCredit() string {
	if x != nil {
		return x.DebitCredit
	}
	return ""
}

func (x *Transaction) GetTransDescription() string {
	if x != nil {
		return x.TransDescription
	}
	return ""
}

func (x *Transaction) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Transaction) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetTransactionList() []*TransactionList {
	if x != nil {
		return x.TransactionList
	}
	return nil
}

type TransactionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId        string `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	DateOn           string `protobuf:"bytes,3,opt,name=dateOn,proto3" json:"dateOn,omitempty"`
	DebitCredit      string `protobuf:"bytes,4,opt,name=debitCredit,proto3" json:"debitCredit,omitempty"`
	TransDescription string `protobuf:"bytes,5,opt,name=transDescription,proto3" json:"transDescription,omitempty"`
	Details          string `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
	Amount           int32  `protobuf:"varint,7,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionList) Reset() {
	*x = TransactionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_uibackend_uibackend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionList) ProtoMessage() {}

func (x *TransactionList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_uibackend_uibackend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionList.ProtoReflect.Descriptor instead.
func (*TransactionList) Descriptor() ([]byte, []int) {
	return file_protos_uibackend_uibackend_proto_rawDescGZIP(), []int{6}
}

func (x *TransactionList) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransactionList) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *TransactionList) GetDateOn() string {
	if x != nil {
		return x.DateOn
	}
	return ""
}

func (x *TransactionList) GetDebitCredit() string {
	if x != nil {
		return x.DebitCredit
	}
	return ""
}

func (x *TransactionList) GetTransDescription() string {
	if x != nil {
		return x.TransDescription
	}
	return ""
}

func (x *TransactionList) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *TransactionList) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_protos_uibackend_uibackend_proto protoreflect.FileDescriptor

var file_protos_uibackend_uibackend_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x5f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xef,
	0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x4f, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42,
	0x79, 0x22, 0x57, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9e, 0x02, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x62, 0x69, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x62, 0x69, 0x74, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x49, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x5f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x62, 0x69, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x62, 0x69,
	0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbe, 0x01, 0x0a, 0x0c, 0x55, 0x69, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x53, 0x76, 0x63, 0x12, 0x50, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x69, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x69,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x5f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x69, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x71, 0x72, 0x61, 0x6d, 0x2f, 0x49, 0x47, 0x2d, 0x54,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x69, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_uibackend_uibackend_proto_rawDescOnce sync.Once
	file_protos_uibackend_uibackend_proto_rawDescData = file_protos_uibackend_uibackend_proto_rawDesc
)

func file_protos_uibackend_uibackend_proto_rawDescGZIP() []byte {
	file_protos_uibackend_uibackend_proto_rawDescOnce.Do(func() {
		file_protos_uibackend_uibackend_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_uibackend_uibackend_proto_rawDescData)
	})
	return file_protos_uibackend_uibackend_proto_rawDescData
}

var file_protos_uibackend_uibackend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_uibackend_uibackend_proto_goTypes = []interface{}{
	(*ProductsRequest)(nil),      // 0: demo_uibackend.ProductsRequest
	(*ProductsResponse)(nil),     // 1: demo_uibackend.ProductsResponse
	(*Product)(nil),              // 2: demo_uibackend.Product
	(*TransactionsRequest)(nil),  // 3: demo_uibackend.TransactionsRequest
	(*TransactionsResponse)(nil), // 4: demo_uibackend.TransactionsResponse
	(*Transaction)(nil),          // 5: demo_uibackend.Transaction
	(*TransactionList)(nil),      // 6: demo_uibackend.TransactionList
}
var file_protos_uibackend_uibackend_proto_depIdxs = []int32{
	2, // 0: demo_uibackend.ProductsResponse.products:type_name -> demo_uibackend.Product
	5, // 1: demo_uibackend.TransactionsResponse.transactions:type_name -> demo_uibackend.Transaction
	6, // 2: demo_uibackend.Transaction.transactionList:type_name -> demo_uibackend.TransactionList
	0, // 3: demo_uibackend.UibackendSvc.getProducts:input_type -> demo_uibackend.ProductsRequest
	3, // 4: demo_uibackend.UibackendSvc.getTransactions:input_type -> demo_uibackend.TransactionsRequest
	1, // 5: demo_uibackend.UibackendSvc.getProducts:output_type -> demo_uibackend.ProductsResponse
	4, // 6: demo_uibackend.UibackendSvc.getTransactions:output_type -> demo_uibackend.TransactionsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_uibackend_uibackend_proto_init() }
func file_protos_uibackend_uibackend_proto_init() {
	if File_protos_uibackend_uibackend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_uibackend_uibackend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_uibackend_uibackend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_uibackend_uibackend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_uibackend_uibackend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_uibackend_uibackend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_uibackend_uibackend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_uibackend_uibackend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_uibackend_uibackend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_uibackend_uibackend_proto_goTypes,
		DependencyIndexes: file_protos_uibackend_uibackend_proto_depIdxs,
		MessageInfos:      file_protos_uibackend_uibackend_proto_msgTypes,
	}.Build()
	File_protos_uibackend_uibackend_proto = out.File
	file_protos_uibackend_uibackend_proto_rawDesc = nil
	file_protos_uibackend_uibackend_proto_goTypes = nil
	file_protos_uibackend_uibackend_proto_depIdxs = nil
}
