// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: api.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PagedRequestOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cursor can be passed to retrieve the next page of results keyed by the cursor
	Cursor int64 `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// page_size specifies the number of items to return in the next page
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PagedRequestOptions) Reset() {
	*x = PagedRequestOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedRequestOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedRequestOptions) ProtoMessage() {}

func (x *PagedRequestOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedRequestOptions.ProtoReflect.Descriptor instead.
func (*PagedRequestOptions) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *PagedRequestOptions) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *PagedRequestOptions) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PagedResponseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A cursor that can be provided to retrieve the next page of results
	NextCursor int64 `protobuf:"varint,1,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// Whether or not more results exist
	HasMore bool `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *PagedResponseMetadata) Reset() {
	*x = PagedResponseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedResponseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedResponseMetadata) ProtoMessage() {}

func (x *PagedResponseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedResponseMetadata.ProtoReflect.Descriptor instead.
func (*PagedResponseMetadata) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *PagedResponseMetadata) GetNextCursor() int64 {
	if x != nil {
		return x.NextCursor
	}
	return 0
}

func (x *PagedResponseMetadata) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// *******************
// Requests
// *******************
type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountRequest) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type CreateDealsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deals     []*Deal `protobuf:"bytes,1,rep,name=deals,proto3" json:"deals,omitempty"`
	AccountId int64   `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AllItems  bool    `protobuf:"varint,3,opt,name=all_items,json=allItems,proto3" json:"all_items,omitempty"`
}

func (x *CreateDealsRequest) Reset() {
	*x = CreateDealsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDealsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDealsRequest) ProtoMessage() {}

func (x *CreateDealsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDealsRequest.ProtoReflect.Descriptor instead.
func (*CreateDealsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDealsRequest) GetDeals() []*Deal {
	if x != nil {
		return x.Deals
	}
	return nil
}

func (x *CreateDealsRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateDealsRequest) GetAllItems() bool {
	if x != nil {
		return x.AllItems
	}
	return false
}

type CreateOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders    []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	AccountId int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AllItems  bool     `protobuf:"varint,3,opt,name=all_items,json=allItems,proto3" json:"all_items,omitempty"`
}

func (x *CreateOrdersRequest) Reset() {
	*x = CreateOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrdersRequest) ProtoMessage() {}

func (x *CreateOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrdersRequest.ProtoReflect.Descriptor instead.
func (*CreateOrdersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrdersRequest) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *CreateOrdersRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateOrdersRequest) GetAllItems() bool {
	if x != nil {
		return x.AllItems
	}
	return false
}

type CreatePositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*Position `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
	AccountId int64       `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` //Used to close all open positions in the case there is no positions open
	TimeGMT   int64       `protobuf:"varint,3,opt,name=TimeGMT,proto3" json:"TimeGMT,omitempty"`                      //Used to close all open positions in the case there is no positions open
	AllItems  bool        `protobuf:"varint,4,opt,name=all_items,json=allItems,proto3" json:"all_items,omitempty"`
}

func (x *CreatePositionsRequest) Reset() {
	*x = CreatePositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePositionsRequest) ProtoMessage() {}

func (x *CreatePositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePositionsRequest.ProtoReflect.Descriptor instead.
func (*CreatePositionsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePositionsRequest) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *CreatePositionsRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreatePositionsRequest) GetTimeGMT() int64 {
	if x != nil {
		return x.TimeGMT
	}
	return 0
}

func (x *CreatePositionsRequest) GetAllItems() bool {
	if x != nil {
		return x.AllItems
	}
	return false
}

type CreateTradeTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId        int64             `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	CreationOrder    int64             `protobuf:"varint,2,opt,name=creation_order,json=creationOrder,proto3" json:"creation_order,omitempty"`
	TradeTransaction *TradeTransaction `protobuf:"bytes,4,opt,name=trade_transaction,json=tradeTransaction,proto3" json:"trade_transaction,omitempty"`
	TradeRequest     *TradeRequest     `protobuf:"bytes,5,opt,name=trade_request,json=tradeRequest,proto3" json:"trade_request,omitempty"`
	TradeResult      *TradeResult      `protobuf:"bytes,6,opt,name=trade_result,json=tradeResult,proto3" json:"trade_result,omitempty"`
	TimeGMT          int64             `protobuf:"varint,7,opt,name=TimeGMT,proto3" json:"TimeGMT,omitempty"`
}

func (x *CreateTradeTransactionRequest) Reset() {
	*x = CreateTradeTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeTransactionRequest) ProtoMessage() {}

func (x *CreateTradeTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTradeTransactionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *CreateTradeTransactionRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateTradeTransactionRequest) GetCreationOrder() int64 {
	if x != nil {
		return x.CreationOrder
	}
	return 0
}

func (x *CreateTradeTransactionRequest) GetTradeTransaction() *TradeTransaction {
	if x != nil {
		return x.TradeTransaction
	}
	return nil
}

func (x *CreateTradeTransactionRequest) GetTradeRequest() *TradeRequest {
	if x != nil {
		return x.TradeRequest
	}
	return nil
}

func (x *CreateTradeTransactionRequest) GetTradeResult() *TradeResult {
	if x != nil {
		return x.TradeResult
	}
	return nil
}

func (x *CreateTradeTransactionRequest) GetTimeGMT() int64 {
	if x != nil {
		return x.TimeGMT
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x65, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x74, 0x72, 0x61, 0x64, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x13, 0x50, 0x61, 0x67, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x53, 0x0a, 0x15, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x49, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x64, 0x65,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x52,
	0x05, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x47,
	0x4d, 0x54, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x47, 0x4d,
	0x54, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xd1,
	0x02, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65,
	0x47, 0x4d, 0x54, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x47,
	0x4d, 0x54, 0x32, 0xb0, 0x03, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x61, 0x6c, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x51, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x5f, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2d,
	0x62, 0x65, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(*PagedRequestOptions)(nil),           // 0: tickerbeats.v3.PagedRequestOptions
	(*PagedResponseMetadata)(nil),         // 1: tickerbeats.v3.PagedResponseMetadata
	(*CreateAccountRequest)(nil),          // 2: tickerbeats.v3.CreateAccountRequest
	(*CreateDealsRequest)(nil),            // 3: tickerbeats.v3.CreateDealsRequest
	(*CreateOrdersRequest)(nil),           // 4: tickerbeats.v3.CreateOrdersRequest
	(*CreatePositionsRequest)(nil),        // 5: tickerbeats.v3.CreatePositionsRequest
	(*CreateTradeTransactionRequest)(nil), // 6: tickerbeats.v3.CreateTradeTransactionRequest
	(*Account)(nil),                       // 7: tickerbeats.v3.Account
	(*Deal)(nil),                          // 8: tickerbeats.v3.Deal
	(*Order)(nil),                         // 9: tickerbeats.v3.Order
	(*Position)(nil),                      // 10: tickerbeats.v3.Position
	(*TradeTransaction)(nil),              // 11: tickerbeats.v3.TradeTransaction
	(*TradeRequest)(nil),                  // 12: tickerbeats.v3.TradeRequest
	(*TradeResult)(nil),                   // 13: tickerbeats.v3.TradeResult
	(*emptypb.Empty)(nil),                 // 14: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	7,  // 0: tickerbeats.v3.CreateAccountRequest.account:type_name -> tickerbeats.v3.Account
	8,  // 1: tickerbeats.v3.CreateDealsRequest.deals:type_name -> tickerbeats.v3.Deal
	9,  // 2: tickerbeats.v3.CreateOrdersRequest.orders:type_name -> tickerbeats.v3.Order
	10, // 3: tickerbeats.v3.CreatePositionsRequest.positions:type_name -> tickerbeats.v3.Position
	11, // 4: tickerbeats.v3.CreateTradeTransactionRequest.trade_transaction:type_name -> tickerbeats.v3.TradeTransaction
	12, // 5: tickerbeats.v3.CreateTradeTransactionRequest.trade_request:type_name -> tickerbeats.v3.TradeRequest
	13, // 6: tickerbeats.v3.CreateTradeTransactionRequest.trade_result:type_name -> tickerbeats.v3.TradeResult
	2,  // 7: tickerbeats.v3.TransactionsService.CreateAccount:input_type -> tickerbeats.v3.CreateAccountRequest
	3,  // 8: tickerbeats.v3.TransactionsService.CreateDeals:input_type -> tickerbeats.v3.CreateDealsRequest
	4,  // 9: tickerbeats.v3.TransactionsService.CreateOrders:input_type -> tickerbeats.v3.CreateOrdersRequest
	5,  // 10: tickerbeats.v3.TransactionsService.CreatePositions:input_type -> tickerbeats.v3.CreatePositionsRequest
	6,  // 11: tickerbeats.v3.TransactionsService.CreateTradeTransaction:input_type -> tickerbeats.v3.CreateTradeTransactionRequest
	14, // 12: tickerbeats.v3.TransactionsService.CreateAccount:output_type -> google.protobuf.Empty
	14, // 13: tickerbeats.v3.TransactionsService.CreateDeals:output_type -> google.protobuf.Empty
	14, // 14: tickerbeats.v3.TransactionsService.CreateOrders:output_type -> google.protobuf.Empty
	14, // 15: tickerbeats.v3.TransactionsService.CreatePositions:output_type -> google.protobuf.Empty
	14, // 16: tickerbeats.v3.TransactionsService.CreateTradeTransaction:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_account_proto_init()
	file_order_proto_init()
	file_deal_proto_init()
	file_position_proto_init()
	file_tradetransaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedRequestOptions); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedResponseMetadata); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDealsRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrdersRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePositionsRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeTransactionRequest); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}