// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: position.proto

package v3

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

// PositionType: Position Type
type PositionType int32

const (
	PositionType_POSITION_TYPE_BUY  PositionType = 0 //Buy
	PositionType_POSITION_TYPE_SELL PositionType = 1 //Sell
)

// Enum value maps for PositionType.
var (
	PositionType_name = map[int32]string{
		0: "POSITION_TYPE_BUY",
		1: "POSITION_TYPE_SELL",
	}
	PositionType_value = map[string]int32{
		"POSITION_TYPE_BUY":  0,
		"POSITION_TYPE_SELL": 1,
	}
)

func (x PositionType) Enum() *PositionType {
	p := new(PositionType)
	*p = x
	return p
}

func (x PositionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PositionType) Descriptor() protoreflect.EnumDescriptor {
	return file_position_proto_enumTypes[0].Descriptor()
}

func (PositionType) Type() protoreflect.EnumType {
	return &file_position_proto_enumTypes[0]
}

func (x PositionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PositionType.Descriptor instead.
func (PositionType) EnumDescriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

type PositionReason int32

const (
	PositionReason_POSITION_REASON_CLIENT PositionReason = 0 //The position was opened as a result of activation of an order placed from a desktop terminal
	PositionReason_POSITION_REASON_MOBILE PositionReason = 1 //The position was opened as a result of activation of an order placed from a mobile application
	PositionReason_POSITION_REASON_WEB    PositionReason = 2 //The position was opened as a result of activation of an order placed from the web platform
	PositionReason_POSITION_REASON_EXPERT PositionReason = 3 //The position was opened as a result of activation of an order placed from an MQL5 program, i.e. an Expert Advisor or a script
)

// Enum value maps for PositionReason.
var (
	PositionReason_name = map[int32]string{
		0: "POSITION_REASON_CLIENT",
		1: "POSITION_REASON_MOBILE",
		2: "POSITION_REASON_WEB",
		3: "POSITION_REASON_EXPERT",
	}
	PositionReason_value = map[string]int32{
		"POSITION_REASON_CLIENT": 0,
		"POSITION_REASON_MOBILE": 1,
		"POSITION_REASON_WEB":    2,
		"POSITION_REASON_EXPERT": 3,
	}
)

func (x PositionReason) Enum() *PositionReason {
	p := new(PositionReason)
	*p = x
	return p
}

func (x PositionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PositionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_position_proto_enumTypes[1].Descriptor()
}

func (PositionReason) Type() protoreflect.EnumType {
	return &file_position_proto_enumTypes[1]
}

func (x PositionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PositionReason.Descriptor instead.
func (PositionReason) EnumDescriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{1}
}

// Position: Execution of trade operations results in the opening of a position
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId     int64          `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	AccountId      int64          `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Ticket         int64          `protobuf:"varint,3,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Symbol         string         `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	PositionTime   int64          `protobuf:"varint,5,opt,name=position_time,json=positionTime,proto3" json:"position_time,omitempty"`
	PositionType   PositionType   `protobuf:"varint,6,opt,name=position_type,json=positionType,proto3,enum=tickerbeats.v3.PositionType" json:"position_type,omitempty"`
	Volume         float64        `protobuf:"fixed64,7,opt,name=volume,proto3" json:"volume,omitempty"`
	PriceOpen      float64        `protobuf:"fixed64,8,opt,name=price_open,json=priceOpen,proto3" json:"price_open,omitempty"`
	StopLoss       float64        `protobuf:"fixed64,9,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`
	TakeProfit     float64        `protobuf:"fixed64,10,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`
	PriceCurrent   float64        `protobuf:"fixed64,11,opt,name=price_current,json=priceCurrent,proto3" json:"price_current,omitempty"`
	Commission     float64        `protobuf:"fixed64,12,opt,name=commission,proto3" json:"commission,omitempty"`
	Swap           float64        `protobuf:"fixed64,13,opt,name=swap,proto3" json:"swap,omitempty"`
	Profit         float64        `protobuf:"fixed64,14,opt,name=profit,proto3" json:"profit,omitempty"`
	Identifier     int64          `protobuf:"varint,15,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Comment        string         `protobuf:"bytes,16,opt,name=comment,proto3" json:"comment,omitempty"`
	Created        int64          `protobuf:"varint,17,opt,name=created,proto3" json:"created,omitempty"`
	Updated        int64          `protobuf:"varint,18,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted        int64          `protobuf:"varint,19,opt,name=deleted,proto3" json:"deleted,omitempty"`
	PositionUpdate int64          `protobuf:"varint,20,opt,name=position_update,json=positionUpdate,proto3" json:"position_update,omitempty"`
	Reason         PositionReason `protobuf:"varint,21,opt,name=reason,proto3,enum=tickerbeats.v3.PositionReason" json:"reason,omitempty"`
	ExternalId     string         `protobuf:"bytes,22,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Magic          int64          `protobuf:"varint,23,opt,name=magic,proto3" json:"magic,omitempty"`
	TimeGMTOffset  int64          `protobuf:"varint,24,opt,name=TimeGMTOffset,proto3" json:"TimeGMTOffset,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *Position) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Position) GetTicket() int64 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *Position) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Position) GetPositionTime() int64 {
	if x != nil {
		return x.PositionTime
	}
	return 0
}

func (x *Position) GetPositionType() PositionType {
	if x != nil {
		return x.PositionType
	}
	return PositionType_POSITION_TYPE_BUY
}

func (x *Position) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Position) GetPriceOpen() float64 {
	if x != nil {
		return x.PriceOpen
	}
	return 0
}

func (x *Position) GetStopLoss() float64 {
	if x != nil {
		return x.StopLoss
	}
	return 0
}

func (x *Position) GetTakeProfit() float64 {
	if x != nil {
		return x.TakeProfit
	}
	return 0
}

func (x *Position) GetPriceCurrent() float64 {
	if x != nil {
		return x.PriceCurrent
	}
	return 0
}

func (x *Position) GetCommission() float64 {
	if x != nil {
		return x.Commission
	}
	return 0
}

func (x *Position) GetSwap() float64 {
	if x != nil {
		return x.Swap
	}
	return 0
}

func (x *Position) GetProfit() float64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *Position) GetIdentifier() int64 {
	if x != nil {
		return x.Identifier
	}
	return 0
}

func (x *Position) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Position) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Position) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Position) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *Position) GetPositionUpdate() int64 {
	if x != nil {
		return x.PositionUpdate
	}
	return 0
}

func (x *Position) GetReason() PositionReason {
	if x != nil {
		return x.Reason
	}
	return PositionReason_POSITION_REASON_CLIENT
}

func (x *Position) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Position) GetMagic() int64 {
	if x != nil {
		return x.Magic
	}
	return 0
}

func (x *Position) GetTimeGMTOffset() int64 {
	if x != nil {
		return x.TimeGMTOffset
	}
	return 0
}

var File_position_proto protoreflect.FileDescriptor

var file_position_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x22, 0x8e, 0x06, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x6b,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x77, 0x61, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73,
	0x77, 0x61, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x12, 0x24, 0x0a, 0x0d, 0x54,
	0x69, 0x6d, 0x65, 0x47, 0x4d, 0x54, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x47, 0x4d, 0x54, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x2a, 0x3d, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01,
	0x2a, 0x7d, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x57, 0x45,
	0x42, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x54, 0x10, 0x03, 0x42,
	0x11, 0x5a, 0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2f,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_position_proto_rawDescOnce sync.Once
	file_position_proto_rawDescData = file_position_proto_rawDesc
)

func file_position_proto_rawDescGZIP() []byte {
	file_position_proto_rawDescOnce.Do(func() {
		file_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_position_proto_rawDescData)
	})
	return file_position_proto_rawDescData
}

var file_position_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_position_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_position_proto_goTypes = []interface{}{
	(PositionType)(0),   // 0: tickerbeats.v3.PositionType
	(PositionReason)(0), // 1: tickerbeats.v3.PositionReason
	(*Position)(nil),    // 2: tickerbeats.v3.Position
}
var file_position_proto_depIdxs = []int32{
	0, // 0: tickerbeats.v3.Position.position_type:type_name -> tickerbeats.v3.PositionType
	1, // 1: tickerbeats.v3.Position.reason:type_name -> tickerbeats.v3.PositionReason
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_position_proto_init() }
func file_position_proto_init() {
	if File_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
			RawDescriptor: file_position_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_position_proto_goTypes,
		DependencyIndexes: file_position_proto_depIdxs,
		EnumInfos:         file_position_proto_enumTypes,
		MessageInfos:      file_position_proto_msgTypes,
	}.Build()
	File_position_proto = out.File
	file_position_proto_rawDesc = nil
	file_position_proto_goTypes = nil
	file_position_proto_depIdxs = nil
}
