// Copyright 2021 Sander Ruscigno
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: signal.proto

package tickerbeats_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SignalType: type of the signal
type SignalType int32

const (
	SignalType_SIGINAL_TYPE_ORDER    SignalType = 0 // Order beats
	SignalType_SIGINAL_TYPE_POSITION SignalType = 1 // Position beats
)

// Enum value maps for SignalType.
var (
	SignalType_name = map[int32]string{
		0: "SIGINAL_TYPE_ORDER",
		1: "SIGINAL_TYPE_POSITION",
	}
	SignalType_value = map[string]int32{
		"SIGINAL_TYPE_ORDER":    0,
		"SIGINAL_TYPE_POSITION": 1,
	}
)

func (x SignalType) Enum() *SignalType {
	p := new(SignalType)
	*p = x
	return p
}

func (x SignalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignalType) Descriptor() protoreflect.EnumDescriptor {
	return file_signal_proto_enumTypes[0].Descriptor()
}

func (SignalType) Type() protoreflect.EnumType {
	return &file_signal_proto_enumTypes[0]
}

func (x SignalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignalType.Descriptor instead.
func (SignalType) EnumDescriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalId             int64   `protobuf:"varint,1,opt,name=SignalId,proto3" json:"SignalId,omitempty"`
	SourceAccountId      int64   `protobuf:"varint,2,opt,name=SourceAccountId,proto3" json:"SourceAccountId,omitempty"`
	DestinationAccountId int64   `protobuf:"varint,3,opt,name=DestinationAccountId,proto3" json:"DestinationAccountId,omitempty"`
	Active               bool    `protobuf:"varint,4,opt,name=Active,proto3" json:"Active,omitempty"`
	MaxDeposit           int32   `protobuf:"varint,5,opt,name=MaxDeposit,proto3" json:"MaxDeposit,omitempty"`
	StopIfLessThan       int32   `protobuf:"varint,6,opt,name=StopIfLessThan,proto3" json:"StopIfLessThan,omitempty"`
	MaxSpread            float64 `protobuf:"fixed64,7,opt,name=MaxSpread,proto3" json:"MaxSpread,omitempty"`
	MinutesToExpire      int32   `protobuf:"varint,8,opt,name=MinutesToExpire,proto3" json:"MinutesToExpire,omitempty"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

func (x *Signal) GetSignalId() int64 {
	if x != nil {
		return x.SignalId
	}
	return 0
}

func (x *Signal) GetSourceAccountId() int64 {
	if x != nil {
		return x.SourceAccountId
	}
	return 0
}

func (x *Signal) GetDestinationAccountId() int64 {
	if x != nil {
		return x.DestinationAccountId
	}
	return 0
}

func (x *Signal) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Signal) GetMaxDeposit() int32 {
	if x != nil {
		return x.MaxDeposit
	}
	return 0
}

func (x *Signal) GetStopIfLessThan() int32 {
	if x != nil {
		return x.StopIfLessThan
	}
	return 0
}

func (x *Signal) GetMaxSpread() float64 {
	if x != nil {
		return x.MaxSpread
	}
	return 0
}

func (x *Signal) GetMinutesToExpire() int32 {
	if x != nil {
		return x.MinutesToExpire
	}
	return 0
}

type SignalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalResultId       int64                  `protobuf:"varint,1,opt,name=SignalResultId,proto3" json:"SignalResultId,omitempty"`
	SourceAccountId      int64                  `protobuf:"varint,2,opt,name=SourceAccountId,proto3" json:"SourceAccountId,omitempty"`
	DestinationAccountId int64                  `protobuf:"varint,3,opt,name=DestinationAccountId,proto3" json:"DestinationAccountId,omitempty"`
	SourceBeatsId        int64                  `protobuf:"varint,4,opt,name=SourceBeatsId,proto3" json:"SourceBeatsId,omitempty"`
	DestinationBeatsId   int64                  `protobuf:"varint,5,opt,name=DestinationBeatsId,proto3" json:"DestinationBeatsId,omitempty"`
	ExternalId           int64                  `protobuf:"varint,6,opt,name=ExternalId,proto3" json:"ExternalId,omitempty"`
	SentTime             *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=SentTime,proto3" json:"SentTime,omitempty"`
	ConfirmationTime     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=ConfirmationTime,proto3" json:"ConfirmationTime,omitempty"`
	Created              *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=Created,proto3" json:"Created,omitempty"`
	Updated              *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=Updated,proto3" json:"Updated,omitempty"`
}

func (x *SignalResult) Reset() {
	*x = SignalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalResult) ProtoMessage() {}

func (x *SignalResult) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalResult.ProtoReflect.Descriptor instead.
func (*SignalResult) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{1}
}

func (x *SignalResult) GetSignalResultId() int64 {
	if x != nil {
		return x.SignalResultId
	}
	return 0
}

func (x *SignalResult) GetSourceAccountId() int64 {
	if x != nil {
		return x.SourceAccountId
	}
	return 0
}

func (x *SignalResult) GetDestinationAccountId() int64 {
	if x != nil {
		return x.DestinationAccountId
	}
	return 0
}

func (x *SignalResult) GetSourceBeatsId() int64 {
	if x != nil {
		return x.SourceBeatsId
	}
	return 0
}

func (x *SignalResult) GetDestinationBeatsId() int64 {
	if x != nil {
		return x.DestinationBeatsId
	}
	return 0
}

func (x *SignalResult) GetExternalId() int64 {
	if x != nil {
		return x.ExternalId
	}
	return 0
}

func (x *SignalResult) GetSentTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SentTime
	}
	return nil
}

func (x *SignalResult) GetConfirmationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ConfirmationTime
	}
	return nil
}

func (x *SignalResult) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *SignalResult) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

var File_signal_proto protoreflect.FileDescriptor

var file_signal_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaa, 0x02, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x4d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x4d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x53, 0x74, 0x6f, 0x70, 0x49, 0x66, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x49, 0x66, 0x4c, 0x65, 0x73, 0x73,
	0x54, 0x68, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x53, 0x70, 0x72, 0x65, 0x61,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4d, 0x61, 0x78, 0x53, 0x70, 0x72, 0x65,
	0x61, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0xf6, 0x03, 0x0a,
	0x0c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x32, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x65, 0x61,
	0x74, 0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x46, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x2a, 0x3f, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53,
	0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signal_proto_rawDescOnce sync.Once
	file_signal_proto_rawDescData = file_signal_proto_rawDesc
)

func file_signal_proto_rawDescGZIP() []byte {
	file_signal_proto_rawDescOnce.Do(func() {
		file_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_signal_proto_rawDescData)
	})
	return file_signal_proto_rawDescData
}

var file_signal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_signal_proto_goTypes = []interface{}{
	(SignalType)(0),               // 0: tickerbeats.v1.SignalType
	(*Signal)(nil),                // 1: tickerbeats.v1.Signal
	(*SignalResult)(nil),          // 2: tickerbeats.v1.SignalResult
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_signal_proto_depIdxs = []int32{
	3, // 0: tickerbeats.v1.SignalResult.SentTime:type_name -> google.protobuf.Timestamp
	3, // 1: tickerbeats.v1.SignalResult.ConfirmationTime:type_name -> google.protobuf.Timestamp
	3, // 2: tickerbeats.v1.SignalResult.Created:type_name -> google.protobuf.Timestamp
	3, // 3: tickerbeats.v1.SignalResult.Updated:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_signal_proto_init() }
func file_signal_proto_init() {
	if File_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
		file_signal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalResult); i {
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
			RawDescriptor: file_signal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signal_proto_goTypes,
		DependencyIndexes: file_signal_proto_depIdxs,
		EnumInfos:         file_signal_proto_enumTypes,
		MessageInfos:      file_signal_proto_msgTypes,
	}.Build()
	File_signal_proto = out.File
	file_signal_proto_rawDesc = nil
	file_signal_proto_goTypes = nil
	file_signal_proto_depIdxs = nil
}
