// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: order.proto

package v1

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

// OrderType: Order type
type OrderType int32

const (
	OrderType_ORDER_TYPE_BUY             OrderType = 0 //Market Buy order
	OrderType_ORDER_TYPE_SELL            OrderType = 1 //Market Sell order
	OrderType_ORDER_TYPE_BUY_LIMIT       OrderType = 2 //Buy Limit pending order
	OrderType_ORDER_TYPE_SELL_LIMIT      OrderType = 3 //Sell Limit pending order
	OrderType_ORDER_TYPE_BUY_STOP        OrderType = 4 //Buy Stop pending order
	OrderType_ORDER_TYPE_SELL_STOP       OrderType = 5 //Sell Stop pending order
	OrderType_ORDER_TYPE_BUY_STOP_LIMIT  OrderType = 6 //Upon reaching the order price, a pending Buy Limit order is placed at the StopLimit price
	OrderType_ORDER_TYPE_SELL_STOP_LIMIT OrderType = 7 //Upon reaching the order price, a pending Sell Limit order is placed at the StopLimit price
	OrderType_ORDER_TYPE_CLOSE_BY        OrderType = 8 //Order to close a position by an opposite one
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "ORDER_TYPE_BUY",
		1: "ORDER_TYPE_SELL",
		2: "ORDER_TYPE_BUY_LIMIT",
		3: "ORDER_TYPE_SELL_LIMIT",
		4: "ORDER_TYPE_BUY_STOP",
		5: "ORDER_TYPE_SELL_STOP",
		6: "ORDER_TYPE_BUY_STOP_LIMIT",
		7: "ORDER_TYPE_SELL_STOP_LIMIT",
		8: "ORDER_TYPE_CLOSE_BY",
	}
	OrderType_value = map[string]int32{
		"ORDER_TYPE_BUY":             0,
		"ORDER_TYPE_SELL":            1,
		"ORDER_TYPE_BUY_LIMIT":       2,
		"ORDER_TYPE_SELL_LIMIT":      3,
		"ORDER_TYPE_BUY_STOP":        4,
		"ORDER_TYPE_SELL_STOP":       5,
		"ORDER_TYPE_BUY_STOP_LIMIT":  6,
		"ORDER_TYPE_SELL_STOP_LIMIT": 7,
		"ORDER_TYPE_CLOSE_BY":        8,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

// OrderType: Order state
type OrderState int32

const (
	OrderState_ORDER_STATE_STARTED        OrderState = 0 //Order checked, but not yet accepted by broker
	OrderState_ORDER_STATE_PLACED         OrderState = 1 //Order accepted
	OrderState_ORDER_STATE_CANCELED       OrderState = 2 //Order canceled by client
	OrderState_ORDER_STATE_PARTIAL        OrderState = 3 //Order partially executed
	OrderState_ORDER_STATE_FILLED         OrderState = 4 //Order fully executed
	OrderState_ORDER_STATE_REJECTED       OrderState = 5 //Order rejected
	OrderState_ORDER_STATE_EXPIRED        OrderState = 6 //Order expired
	OrderState_ORDER_STATE_REQUEST_ADD    OrderState = 7 //Order is being registered (placing to the trading system)
	OrderState_ORDER_STATE_REQUEST_MODIFY OrderState = 8 //Order is being modified (changing its parameters)
	OrderState_ORDER_STATE_REQUEST_CANCEL OrderState = 9 //Order is being deleted (deleting from the trading system)
)

// Enum value maps for OrderState.
var (
	OrderState_name = map[int32]string{
		0: "ORDER_STATE_STARTED",
		1: "ORDER_STATE_PLACED",
		2: "ORDER_STATE_CANCELED",
		3: "ORDER_STATE_PARTIAL",
		4: "ORDER_STATE_FILLED",
		5: "ORDER_STATE_REJECTED",
		6: "ORDER_STATE_EXPIRED",
		7: "ORDER_STATE_REQUEST_ADD",
		8: "ORDER_STATE_REQUEST_MODIFY",
		9: "ORDER_STATE_REQUEST_CANCEL",
	}
	OrderState_value = map[string]int32{
		"ORDER_STATE_STARTED":        0,
		"ORDER_STATE_PLACED":         1,
		"ORDER_STATE_CANCELED":       2,
		"ORDER_STATE_PARTIAL":        3,
		"ORDER_STATE_FILLED":         4,
		"ORDER_STATE_REJECTED":       5,
		"ORDER_STATE_EXPIRED":        6,
		"ORDER_STATE_REQUEST_ADD":    7,
		"ORDER_STATE_REQUEST_MODIFY": 8,
		"ORDER_STATE_REQUEST_CANCEL": 9,
	}
)

func (x OrderState) Enum() *OrderState {
	p := new(OrderState)
	*p = x
	return p
}

func (x OrderState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderState) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[1].Descriptor()
}

func (OrderState) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[1]
}

func (x OrderState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderState.Descriptor instead.
func (OrderState) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

// OrderFillingType: Order filling type
type OrderFillingType int32

const (
	OrderFillingType_ORDER_FILLING_FOK    OrderFillingType = 0 //This filling policy means that an order can be filled only in the specified amount.
	OrderFillingType_ORDER_FILLING_IOC    OrderFillingType = 1 //This mode means that a trader agrees to execute a deal with the volume maximally available in the market within that indicated in the order.
	OrderFillingType_ORDER_FILLING_RETURN OrderFillingType = 2 //This policy is used only for market orders (ORDER_TYPE_BUY and ORDER_TYPE_SELL), limit and stop limit orders (ORDER_TYPE_BUY_LIMIT, ORDER_TYPE_SELL_LIMIT, ORDER_TYPE_BUY_STOP_LIMIT and ORDER_TYPE_SELL_STOP_LIMIT ) and only for the symbols with Market or Exchange execution.
)

// Enum value maps for OrderFillingType.
var (
	OrderFillingType_name = map[int32]string{
		0: "ORDER_FILLING_FOK",
		1: "ORDER_FILLING_IOC",
		2: "ORDER_FILLING_RETURN",
	}
	OrderFillingType_value = map[string]int32{
		"ORDER_FILLING_FOK":    0,
		"ORDER_FILLING_IOC":    1,
		"ORDER_FILLING_RETURN": 2,
	}
)

func (x OrderFillingType) Enum() *OrderFillingType {
	p := new(OrderFillingType)
	*p = x
	return p
}

func (x OrderFillingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderFillingType) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[2].Descriptor()
}

func (OrderFillingType) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[2]
}

func (x OrderFillingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderFillingType.Descriptor instead.
func (OrderFillingType) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

// OrderTypeTime: Order lifetime
type OrderTypeTime int32

const (
	OrderTypeTime_ORDER_TIME_GTC           OrderTypeTime = 0 //Good till cancel order
	OrderTypeTime_ORDER_TIME_DAY           OrderTypeTime = 1 //Good till current trade day order
	OrderTypeTime_ORDER_TIME_SPECIFIED     OrderTypeTime = 2 //Good till expired order
	OrderTypeTime_ORDER_TIME_SPECIFIED_DAY OrderTypeTime = 3 //The order will be effective till 23:59:59 of the specified day.
)

// Enum value maps for OrderTypeTime.
var (
	OrderTypeTime_name = map[int32]string{
		0: "ORDER_TIME_GTC",
		1: "ORDER_TIME_DAY",
		2: "ORDER_TIME_SPECIFIED",
		3: "ORDER_TIME_SPECIFIED_DAY",
	}
	OrderTypeTime_value = map[string]int32{
		"ORDER_TIME_GTC":           0,
		"ORDER_TIME_DAY":           1,
		"ORDER_TIME_SPECIFIED":     2,
		"ORDER_TIME_SPECIFIED_DAY": 3,
	}
)

func (x OrderTypeTime) Enum() *OrderTypeTime {
	p := new(OrderTypeTime)
	*p = x
	return p
}

func (x OrderTypeTime) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderTypeTime) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[3].Descriptor()
}

func (OrderTypeTime) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[3]
}

func (x OrderTypeTime) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderTypeTime.Descriptor instead.
func (OrderTypeTime) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

// OrderReason: The reason for order placing. An order can be placed by an MQL5 program, from a mobile application, as a result of StopOut, etc.
type OrderReason int32

const (
	OrderReason_ORDER_REASON_CLIENT OrderReason = 0 //The order was placed from a desktop terminal
	OrderReason_ORDER_REASON_MOBILE OrderReason = 1 //The order was placed from a mobile application
	OrderReason_ORDER_REASON_WEB    OrderReason = 2 //The order was placed from a web platform
	OrderReason_ORDER_REASON_EXPERT OrderReason = 3 //The order was placed from an MQL5-program, i.e. by an Expert Advisor or a script
	OrderReason_ORDER_REASON_SL     OrderReason = 4 //The order was placed as a result of Stop Loss activation
	OrderReason_ORDER_REASON_TP     OrderReason = 5 //The order was placed as a result of Take Profit activation
	OrderReason_ORDER_REASON_SO     OrderReason = 6 //The order was placed as a result of the Stop Out event
)

// Enum value maps for OrderReason.
var (
	OrderReason_name = map[int32]string{
		0: "ORDER_REASON_CLIENT",
		1: "ORDER_REASON_MOBILE",
		2: "ORDER_REASON_WEB",
		3: "ORDER_REASON_EXPERT",
		4: "ORDER_REASON_SL",
		5: "ORDER_REASON_TP",
		6: "ORDER_REASON_SO",
	}
	OrderReason_value = map[string]int32{
		"ORDER_REASON_CLIENT": 0,
		"ORDER_REASON_MOBILE": 1,
		"ORDER_REASON_WEB":    2,
		"ORDER_REASON_EXPERT": 3,
		"ORDER_REASON_SL":     4,
		"ORDER_REASON_TP":     5,
		"ORDER_REASON_SO":     6,
	}
)

func (x OrderReason) Enum() *OrderReason {
	p := new(OrderReason)
	*p = x
	return p
}

func (x OrderReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderReason) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[4].Descriptor()
}

func (OrderReason) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[4]
}

func (x OrderReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderReason.Descriptor instead.
func (OrderReason) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

// Order represents an trade order
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId        int64            `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	AccountId      int64            `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Ticket         int64            `protobuf:"varint,3,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Symbol         string           `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	TimeSetup      int64            `protobuf:"varint,5,opt,name=time_setup,json=timeSetup,proto3" json:"time_setup,omitempty"`
	OrderType      OrderType        `protobuf:"varint,6,opt,name=order_type,json=orderType,proto3,enum=tickerbeats.v1.OrderType" json:"order_type,omitempty"`
	State          OrderState       `protobuf:"varint,7,opt,name=state,proto3,enum=tickerbeats.v1.OrderState" json:"state,omitempty"`
	TimeExpiration int64            `protobuf:"varint,8,opt,name=time_expiration,json=timeExpiration,proto3" json:"time_expiration,omitempty"`
	TimeDone       int64            `protobuf:"varint,9,opt,name=time_done,json=timeDone,proto3" json:"time_done,omitempty"`
	TypeFilling    OrderFillingType `protobuf:"varint,10,opt,name=type_filling,json=typeFilling,proto3,enum=tickerbeats.v1.OrderFillingType" json:"type_filling,omitempty"`
	TypeTime       OrderTypeTime    `protobuf:"varint,11,opt,name=type_time,json=typeTime,proto3,enum=tickerbeats.v1.OrderTypeTime" json:"type_time,omitempty"`
	Magic          int64            `protobuf:"varint,12,opt,name=magic,proto3" json:"magic,omitempty"`
	PositionId     int64            `protobuf:"varint,13,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Reason         OrderReason      `protobuf:"varint,14,opt,name=reason,proto3,enum=tickerbeats.v1.OrderReason" json:"reason,omitempty"`
	VolumeInitial  float64          `protobuf:"fixed64,15,opt,name=volume_initial,json=volumeInitial,proto3" json:"volume_initial,omitempty"`
	VolumeCurrent  float64          `protobuf:"fixed64,16,opt,name=volume_current,json=volumeCurrent,proto3" json:"volume_current,omitempty"`
	PriceOpen      float64          `protobuf:"fixed64,17,opt,name=price_open,json=priceOpen,proto3" json:"price_open,omitempty"`
	StopLoss       float64          `protobuf:"fixed64,18,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`
	TakeProfit     float64          `protobuf:"fixed64,19,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`
	PriceCurrent   float64          `protobuf:"fixed64,20,opt,name=price_current,json=priceCurrent,proto3" json:"price_current,omitempty"`
	PriceStopLimit float64          `protobuf:"fixed64,21,opt,name=price_stop_limit,json=priceStopLimit,proto3" json:"price_stop_limit,omitempty"`
	Comment        string           `protobuf:"bytes,22,opt,name=comment,proto3" json:"comment,omitempty"`
	ExternalId     string           `protobuf:"bytes,23,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Created        int64            `protobuf:"varint,24,opt,name=created,proto3" json:"created,omitempty"`
	Updated        int64            `protobuf:"varint,25,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted        int64            `protobuf:"varint,26,opt,name=deleted,proto3" json:"deleted,omitempty"`
	PositionById   int64            `protobuf:"varint,27,opt,name=position_by_id,json=positionById,proto3" json:"position_by_id,omitempty"`
	Direction      bool             `protobuf:"varint,28,opt,name=direction,proto3" json:"direction,omitempty"`
	TimeGMTOffset  int64            `protobuf:"varint,29,opt,name=TimeGMTOffset,proto3" json:"TimeGMTOffset,omitempty"`
	ProviderId     string           `protobuf:"bytes,30,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	InternalId     string           `protobuf:"bytes,31,opt,name=internal_id,json=internalId,proto3" json:"internal_id,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Order) GetTicket() int64 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *Order) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Order) GetTimeSetup() int64 {
	if x != nil {
		return x.TimeSetup
	}
	return 0
}

func (x *Order) GetOrderType() OrderType {
	if x != nil {
		return x.OrderType
	}
	return OrderType_ORDER_TYPE_BUY
}

func (x *Order) GetState() OrderState {
	if x != nil {
		return x.State
	}
	return OrderState_ORDER_STATE_STARTED
}

func (x *Order) GetTimeExpiration() int64 {
	if x != nil {
		return x.TimeExpiration
	}
	return 0
}

func (x *Order) GetTimeDone() int64 {
	if x != nil {
		return x.TimeDone
	}
	return 0
}

func (x *Order) GetTypeFilling() OrderFillingType {
	if x != nil {
		return x.TypeFilling
	}
	return OrderFillingType_ORDER_FILLING_FOK
}

func (x *Order) GetTypeTime() OrderTypeTime {
	if x != nil {
		return x.TypeTime
	}
	return OrderTypeTime_ORDER_TIME_GTC
}

func (x *Order) GetMagic() int64 {
	if x != nil {
		return x.Magic
	}
	return 0
}

func (x *Order) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *Order) GetReason() OrderReason {
	if x != nil {
		return x.Reason
	}
	return OrderReason_ORDER_REASON_CLIENT
}

func (x *Order) GetVolumeInitial() float64 {
	if x != nil {
		return x.VolumeInitial
	}
	return 0
}

func (x *Order) GetVolumeCurrent() float64 {
	if x != nil {
		return x.VolumeCurrent
	}
	return 0
}

func (x *Order) GetPriceOpen() float64 {
	if x != nil {
		return x.PriceOpen
	}
	return 0
}

func (x *Order) GetStopLoss() float64 {
	if x != nil {
		return x.StopLoss
	}
	return 0
}

func (x *Order) GetTakeProfit() float64 {
	if x != nil {
		return x.TakeProfit
	}
	return 0
}

func (x *Order) GetPriceCurrent() float64 {
	if x != nil {
		return x.PriceCurrent
	}
	return 0
}

func (x *Order) GetPriceStopLimit() float64 {
	if x != nil {
		return x.PriceStopLimit
	}
	return 0
}

func (x *Order) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Order) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Order) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Order) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Order) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *Order) GetPositionById() int64 {
	if x != nil {
		return x.PositionById
	}
	return 0
}

func (x *Order) GetDirection() bool {
	if x != nil {
		return x.Direction
	}
	return false
}

func (x *Order) GetTimeGMTOffset() int64 {
	if x != nil {
		return x.TimeGMTOffset
	}
	return 0
}

func (x *Order) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *Order) GetInternalId() string {
	if x != nil {
		return x.InternalId
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xde, 0x08,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x12, 0x38, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x6f,
	0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x6f,
	0x6e, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65,
	0x46, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x6b,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x47, 0x4d,
	0x54, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54,
	0x69, 0x6d, 0x65, 0x47, 0x4d, 0x54, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x2a, 0xf4,
	0x01, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45,
	0x4c, 0x4c, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x53, 0x54, 0x4f,
	0x50, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x05, 0x12, 0x1d, 0x0a,
	0x19, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x5f,
	0x53, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x5f,
	0x53, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x5f, 0x42, 0x59, 0x10, 0x08, 0x2a, 0x98, 0x02, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x07,
	0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x10, 0x08,
	0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x09,
	0x2a, 0x5a, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49,
	0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4f, 0x43,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x4c,
	0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x02, 0x2a, 0x6f, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x47, 0x54, 0x43, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x1c, 0x0a, 0x18, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x03, 0x2a, 0xad, 0x01,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x57, 0x45, 0x42, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x54, 0x10, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53,
	0x4c, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x50, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x10, 0x06, 0x42, 0x11, 0x5a,
	0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_order_proto_goTypes = []interface{}{
	(OrderType)(0),        // 0: tickerbeats.v1.OrderType
	(OrderState)(0),       // 1: tickerbeats.v1.OrderState
	(OrderFillingType)(0), // 2: tickerbeats.v1.OrderFillingType
	(OrderTypeTime)(0),    // 3: tickerbeats.v1.OrderTypeTime
	(OrderReason)(0),      // 4: tickerbeats.v1.OrderReason
	(*Order)(nil),         // 5: tickerbeats.v1.Order
}
var file_order_proto_depIdxs = []int32{
	0, // 0: tickerbeats.v1.Order.order_type:type_name -> tickerbeats.v1.OrderType
	1, // 1: tickerbeats.v1.Order.state:type_name -> tickerbeats.v1.OrderState
	2, // 2: tickerbeats.v1.Order.type_filling:type_name -> tickerbeats.v1.OrderFillingType
	3, // 3: tickerbeats.v1.Order.type_time:type_name -> tickerbeats.v1.OrderTypeTime
	4, // 4: tickerbeats.v1.Order.reason:type_name -> tickerbeats.v1.OrderReason
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		EnumInfos:         file_order_proto_enumTypes,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
