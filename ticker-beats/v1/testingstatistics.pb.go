// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: testingstatistics.proto

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

type TestingStatisticsVariables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *TestingStatisticsVariables) Reset() {
	*x = TestingStatisticsVariables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingstatistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestingStatisticsVariables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestingStatisticsVariables) ProtoMessage() {}

func (x *TestingStatisticsVariables) ProtoReflect() protoreflect.Message {
	mi := &file_testingstatistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestingStatisticsVariables.ProtoReflect.Descriptor instead.
func (*TestingStatisticsVariables) Descriptor() ([]byte, []int) {
	return file_testingstatistics_proto_rawDescGZIP(), []int{0}
}

func (x *TestingStatisticsVariables) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestingStatisticsVariables) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TestingStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestingID           int64                         `protobuf:"varint,1,opt,name=TestingID,proto3" json:"TestingID,omitempty"`
	ExpertName          string                        `protobuf:"bytes,2,opt,name=ExpertName,proto3" json:"ExpertName,omitempty"`
	ExpertVersion       string                        `protobuf:"bytes,3,opt,name=ExpertVersion,proto3" json:"ExpertVersion,omitempty"`
	InitialDeposit      float64                       `protobuf:"fixed64,4,opt,name=InitialDeposit,proto3" json:"InitialDeposit,omitempty"`            //The value of the initial deposit
	Withdrawal          float64                       `protobuf:"fixed64,5,opt,name=Withdrawal,proto3" json:"Withdrawal,omitempty"`                    //Money withdrawn from an account
	Profit              float64                       `protobuf:"fixed64,6,opt,name=Profit,proto3" json:"Profit,omitempty"`                            //Net profit after testing, the sum of GROSS_PROFIT and GROSS_LOSS (GROSS_LOSS is always less than or equal to zero)
	GrossProfit         float64                       `protobuf:"fixed64,7,opt,name=GrossProfit,proto3" json:"GrossProfit,omitempty"`                  //Total profit, the sum of all profitable (positive) trades. The value is greater than or equal to zero
	GrossLoss           float64                       `protobuf:"fixed64,8,opt,name=GrossLoss,proto3" json:"GrossLoss,omitempty"`                      //Total loss, the sum of all negative trades. The value is less than or equal to zero
	MaxProfittrade      float64                       `protobuf:"fixed64,9,opt,name=MaxProfittrade,proto3" json:"MaxProfittrade,omitempty"`            //Maximum profit – the largest value of all profitable trades. The value is greater than or equal to zero
	MaxLosstrade        float64                       `protobuf:"fixed64,10,opt,name=MaxLosstrade,proto3" json:"MaxLosstrade,omitempty"`               //Maximum loss – the lowest value of all losing trades. The value is less than or equal to zero
	ConProfitMax        float64                       `protobuf:"fixed64,11,opt,name=ConProfitMax,proto3" json:"ConProfitMax,omitempty"`               //Maximum profit in a series of profitable trades. The value is greater than or equal to zero
	ConProfitMaxTrades  int32                         `protobuf:"varint,12,opt,name=ConProfitMaxTrades,proto3" json:"ConProfitMaxTrades,omitempty"`    //The number of trades that have formed CONPROFITMAX (maximum profit in a series of profitable trades)
	MaxConWins          float64                       `protobuf:"fixed64,13,opt,name=MaxConWins,proto3" json:"MaxConWins,omitempty"`                   //The total profit of the longest series of profitable trades
	MaxConProfitTrades  int32                         `protobuf:"varint,14,opt,name=MaxConProfitTrades,proto3" json:"MaxConProfitTrades,omitempty"`    //The number of trades in the longest series of profitable trades MAX_CONWINS
	ConLossMax          float64                       `protobuf:"fixed64,15,opt,name=ConLossMax,proto3" json:"ConLossMax,omitempty"`                   //Maximum loss in a series of losing trades. The value is less than or equal to zero
	ConLossMaxTrades    int32                         `protobuf:"varint,16,opt,name=ConLossMaxTrades,proto3" json:"ConLossMaxTrades,omitempty"`        //The number of trades that have formed CONLOSSMAX (maximum loss in a series of losing trades)
	MaxConlosses        float64                       `protobuf:"fixed64,17,opt,name=MaxConlosses,proto3" json:"MaxConlosses,omitempty"`               //The total loss of the longest series of losing trades
	MaxConlossTrades    int32                         `protobuf:"varint,18,opt,name=MaxConlossTrades,proto3" json:"MaxConlossTrades,omitempty"`        //The number of trades in the longest series of losing trades MAX_CONLOSSES
	BalanceMin          float64                       `protobuf:"fixed64,19,opt,name=BalanceMin,proto3" json:"BalanceMin,omitempty"`                   //Minimum balance value
	BalanceDD           float64                       `protobuf:"fixed64,20,opt,name=BalanceDD,proto3" json:"BalanceDD,omitempty"`                     //Maximum balance drawdown in monetary terms. In the process of trading, a balance may have numerous drawdowns; here the largest value is taken
	BalanceDDPercent    float64                       `protobuf:"fixed64,21,opt,name=BalanceDDPercent,proto3" json:"BalanceDDPercent,omitempty"`       //Balance drawdown as a percentage that was recorded at the moment of the maximum balance drawdown in monetary terms (BALANCE_DD).
	BalanceDDRelPercent float64                       `protobuf:"fixed64,22,opt,name=BalanceDDRelPercent,proto3" json:"BalanceDDRelPercent,omitempty"` //Maximum balance drawdown as a percentage. In the process of trading, a balance may have numerous drawdowns, for each of which the relative drawdown value in percents is calculated. The greatest value is returned
	BalanceDDRelative   float64                       `protobuf:"fixed64,23,opt,name=BalanceDDRelative,proto3" json:"BalanceDDRelative,omitempty"`     //Balance drawdown in monetary terms that was recorded at the moment of the maximum balance drawdown as a percentage (BALANCE_DDREL_PERCENT).
	EquityMin           float64                       `protobuf:"fixed64,24,opt,name=EquityMin,proto3" json:"EquityMin,omitempty"`                     //Minimum equity value
	EquityDD            float64                       `protobuf:"fixed64,25,opt,name=EquityDD,proto3" json:"EquityDD,omitempty"`                       //Maximum equity drawdown in monetary terms. In the process of trading, numerous drawdowns may appear on the equity; here the largest value is taken
	EquityDDPercent     float64                       `protobuf:"fixed64,26,opt,name=EquityDDPercent,proto3" json:"EquityDDPercent,omitempty"`         //Drawdown in percent that was recorded at the moment of the maximum equity drawdown in monetary terms (EQUITY_DD).
	EquityDDRelPercent  float64                       `protobuf:"fixed64,27,opt,name=EquityDDRelPercent,proto3" json:"EquityDDRelPercent,omitempty"`   //Maximum equity drawdown as a percentage. In the process of trading, an equity may have numerous drawdowns, for each of which the relative drawdown value in percents is calculated. The greatest value is returned
	EquityDDRelative    float64                       `protobuf:"fixed64,28,opt,name=EquityDDRelative,proto3" json:"EquityDDRelative,omitempty"`       //Equity drawdown in monetary terms that was recorded at the moment of the maximum equity drawdown in percent (EQUITY_DDREL_PERCENT).
	ExpectedPayoff      float64                       `protobuf:"fixed64,29,opt,name=ExpectedPayoff,proto3" json:"ExpectedPayoff,omitempty"`           //Expected payoff
	ProfitFactor        float64                       `protobuf:"fixed64,30,opt,name=ProfitFactor,proto3" json:"ProfitFactor,omitempty"`               //Profit factor, equal to the ratio of GROSS_PROFIT/GROSS_LOSS. If GROSS_LOSS=0, the profit factor is equal to DBL_MAX
	RecoveryFactor      float64                       `protobuf:"fixed64,31,opt,name=RecoveryFactor,proto3" json:"RecoveryFactor,omitempty"`           //Recovery factor, equal to the ratio of PROFIT/BALANCE_DD
	SharpeRatio         float64                       `protobuf:"fixed64,32,opt,name=SharpeRatio,proto3" json:"SharpeRatio,omitempty"`                 //Sharpe ratio
	MinMarginlevel      float64                       `protobuf:"fixed64,33,opt,name=MinMarginlevel,proto3" json:"MinMarginlevel,omitempty"`           //Minimum value of the margin level
	CustomOnTester      float64                       `protobuf:"fixed64,34,opt,name=CustomOnTester,proto3" json:"CustomOnTester,omitempty"`           //The value of the calculated custom optimization criterion returned by the OnTester() function
	Deals               int32                         `protobuf:"varint,35,opt,name=Deals,proto3" json:"Deals,omitempty"`                              //The number of deals
	Trades              int32                         `protobuf:"varint,36,opt,name=Trades,proto3" json:"Trades,omitempty"`                            //The number of trades
	ProfitTrades        int32                         `protobuf:"varint,37,opt,name=ProfitTrades,proto3" json:"ProfitTrades,omitempty"`                //Profitable trades
	LossTrades          int32                         `protobuf:"varint,38,opt,name=LossTrades,proto3" json:"LossTrades,omitempty"`                    //Losing trades
	ShortTrades         int32                         `protobuf:"varint,39,opt,name=ShortTrades,proto3" json:"ShortTrades,omitempty"`                  //Short trades
	LongTrades          int32                         `protobuf:"varint,40,opt,name=LongTrades,proto3" json:"LongTrades,omitempty"`                    //Long trades
	ProfitShorttrades   int32                         `protobuf:"varint,41,opt,name=ProfitShorttrades,proto3" json:"ProfitShorttrades,omitempty"`      //Profitable short trades
	ProfitLongtrades    int32                         `protobuf:"varint,42,opt,name=ProfitLongtrades,proto3" json:"ProfitLongtrades,omitempty"`        //Profitable long trades
	ProfitTradesAvgcon  int32                         `protobuf:"varint,43,opt,name=ProfitTradesAvgcon,proto3" json:"ProfitTradesAvgcon,omitempty"`    //Average length of a profitable series of trades
	LossTradesAvgCon    int32                         `protobuf:"varint,44,opt,name=LossTradesAvgCon,proto3" json:"LossTradesAvgCon,omitempty"`        //Average length of a losing series of trades
	MagicNumber         int64                         `protobuf:"varint,45,opt,name=MagicNumber,proto3" json:"MagicNumber,omitempty"`
	Created             int64                         `protobuf:"varint,46,opt,name=Created,proto3" json:"Created,omitempty"`
	Variables           []*TestingStatisticsVariables `protobuf:"bytes,47,rep,name=Variables,proto3" json:"Variables,omitempty"` //Expert variables
}

func (x *TestingStatistics) Reset() {
	*x = TestingStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingstatistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestingStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestingStatistics) ProtoMessage() {}

func (x *TestingStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_testingstatistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestingStatistics.ProtoReflect.Descriptor instead.
func (*TestingStatistics) Descriptor() ([]byte, []int) {
	return file_testingstatistics_proto_rawDescGZIP(), []int{1}
}

func (x *TestingStatistics) GetTestingID() int64 {
	if x != nil {
		return x.TestingID
	}
	return 0
}

func (x *TestingStatistics) GetExpertName() string {
	if x != nil {
		return x.ExpertName
	}
	return ""
}

func (x *TestingStatistics) GetExpertVersion() string {
	if x != nil {
		return x.ExpertVersion
	}
	return ""
}

func (x *TestingStatistics) GetInitialDeposit() float64 {
	if x != nil {
		return x.InitialDeposit
	}
	return 0
}

func (x *TestingStatistics) GetWithdrawal() float64 {
	if x != nil {
		return x.Withdrawal
	}
	return 0
}

func (x *TestingStatistics) GetProfit() float64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *TestingStatistics) GetGrossProfit() float64 {
	if x != nil {
		return x.GrossProfit
	}
	return 0
}

func (x *TestingStatistics) GetGrossLoss() float64 {
	if x != nil {
		return x.GrossLoss
	}
	return 0
}

func (x *TestingStatistics) GetMaxProfittrade() float64 {
	if x != nil {
		return x.MaxProfittrade
	}
	return 0
}

func (x *TestingStatistics) GetMaxLosstrade() float64 {
	if x != nil {
		return x.MaxLosstrade
	}
	return 0
}

func (x *TestingStatistics) GetConProfitMax() float64 {
	if x != nil {
		return x.ConProfitMax
	}
	return 0
}

func (x *TestingStatistics) GetConProfitMaxTrades() int32 {
	if x != nil {
		return x.ConProfitMaxTrades
	}
	return 0
}

func (x *TestingStatistics) GetMaxConWins() float64 {
	if x != nil {
		return x.MaxConWins
	}
	return 0
}

func (x *TestingStatistics) GetMaxConProfitTrades() int32 {
	if x != nil {
		return x.MaxConProfitTrades
	}
	return 0
}

func (x *TestingStatistics) GetConLossMax() float64 {
	if x != nil {
		return x.ConLossMax
	}
	return 0
}

func (x *TestingStatistics) GetConLossMaxTrades() int32 {
	if x != nil {
		return x.ConLossMaxTrades
	}
	return 0
}

func (x *TestingStatistics) GetMaxConlosses() float64 {
	if x != nil {
		return x.MaxConlosses
	}
	return 0
}

func (x *TestingStatistics) GetMaxConlossTrades() int32 {
	if x != nil {
		return x.MaxConlossTrades
	}
	return 0
}

func (x *TestingStatistics) GetBalanceMin() float64 {
	if x != nil {
		return x.BalanceMin
	}
	return 0
}

func (x *TestingStatistics) GetBalanceDD() float64 {
	if x != nil {
		return x.BalanceDD
	}
	return 0
}

func (x *TestingStatistics) GetBalanceDDPercent() float64 {
	if x != nil {
		return x.BalanceDDPercent
	}
	return 0
}

func (x *TestingStatistics) GetBalanceDDRelPercent() float64 {
	if x != nil {
		return x.BalanceDDRelPercent
	}
	return 0
}

func (x *TestingStatistics) GetBalanceDDRelative() float64 {
	if x != nil {
		return x.BalanceDDRelative
	}
	return 0
}

func (x *TestingStatistics) GetEquityMin() float64 {
	if x != nil {
		return x.EquityMin
	}
	return 0
}

func (x *TestingStatistics) GetEquityDD() float64 {
	if x != nil {
		return x.EquityDD
	}
	return 0
}

func (x *TestingStatistics) GetEquityDDPercent() float64 {
	if x != nil {
		return x.EquityDDPercent
	}
	return 0
}

func (x *TestingStatistics) GetEquityDDRelPercent() float64 {
	if x != nil {
		return x.EquityDDRelPercent
	}
	return 0
}

func (x *TestingStatistics) GetEquityDDRelative() float64 {
	if x != nil {
		return x.EquityDDRelative
	}
	return 0
}

func (x *TestingStatistics) GetExpectedPayoff() float64 {
	if x != nil {
		return x.ExpectedPayoff
	}
	return 0
}

func (x *TestingStatistics) GetProfitFactor() float64 {
	if x != nil {
		return x.ProfitFactor
	}
	return 0
}

func (x *TestingStatistics) GetRecoveryFactor() float64 {
	if x != nil {
		return x.RecoveryFactor
	}
	return 0
}

func (x *TestingStatistics) GetSharpeRatio() float64 {
	if x != nil {
		return x.SharpeRatio
	}
	return 0
}

func (x *TestingStatistics) GetMinMarginlevel() float64 {
	if x != nil {
		return x.MinMarginlevel
	}
	return 0
}

func (x *TestingStatistics) GetCustomOnTester() float64 {
	if x != nil {
		return x.CustomOnTester
	}
	return 0
}

func (x *TestingStatistics) GetDeals() int32 {
	if x != nil {
		return x.Deals
	}
	return 0
}

func (x *TestingStatistics) GetTrades() int32 {
	if x != nil {
		return x.Trades
	}
	return 0
}

func (x *TestingStatistics) GetProfitTrades() int32 {
	if x != nil {
		return x.ProfitTrades
	}
	return 0
}

func (x *TestingStatistics) GetLossTrades() int32 {
	if x != nil {
		return x.LossTrades
	}
	return 0
}

func (x *TestingStatistics) GetShortTrades() int32 {
	if x != nil {
		return x.ShortTrades
	}
	return 0
}

func (x *TestingStatistics) GetLongTrades() int32 {
	if x != nil {
		return x.LongTrades
	}
	return 0
}

func (x *TestingStatistics) GetProfitShorttrades() int32 {
	if x != nil {
		return x.ProfitShorttrades
	}
	return 0
}

func (x *TestingStatistics) GetProfitLongtrades() int32 {
	if x != nil {
		return x.ProfitLongtrades
	}
	return 0
}

func (x *TestingStatistics) GetProfitTradesAvgcon() int32 {
	if x != nil {
		return x.ProfitTradesAvgcon
	}
	return 0
}

func (x *TestingStatistics) GetLossTradesAvgCon() int32 {
	if x != nil {
		return x.LossTradesAvgCon
	}
	return 0
}

func (x *TestingStatistics) GetMagicNumber() int64 {
	if x != nil {
		return x.MagicNumber
	}
	return 0
}

func (x *TestingStatistics) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *TestingStatistics) GetVariables() []*TestingStatisticsVariables {
	if x != nil {
		return x.Variables
	}
	return nil
}

var File_testingstatistics_proto protoreflect.FileDescriptor

var file_testingstatistics_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x46, 0x0a, 0x1a, 0x54, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x83, 0x0e, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x65, 0x72, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x47,
	0x72, 0x6f, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x47, 0x72, 0x6f, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x47, 0x72, 0x6f, 0x73, 0x73, 0x4c, 0x6f, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x47, 0x72, 0x6f, 0x73, 0x73, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x4d,
	0x61, 0x78, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x78, 0x4c, 0x6f, 0x73, 0x73, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x4d, 0x61, 0x78, 0x4c, 0x6f,
	0x73, 0x73, 0x74, 0x72, 0x61, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x43,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x2e, 0x0a, 0x12, 0x43,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x4d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4d, 0x61, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x43, 0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4d, 0x61, 0x78, 0x12, 0x2a, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x43, 0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4d, 0x61,
	0x78, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x78, 0x43, 0x6f,
	0x6e, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x4d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x4d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6c, 0x6f, 0x73, 0x73, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6c, 0x6f, 0x73,
	0x73, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x4d, 0x69, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x44, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x44, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x44, 0x44, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x44, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x13, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x44, 0x52, 0x65,
	0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x44, 0x52, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x44,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x44, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x6e, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44, 0x44, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x45,
	0x71, 0x75, 0x69, 0x74, 0x79, 0x44, 0x44, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44, 0x44, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44,
	0x44, 0x52, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x12, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44, 0x44, 0x52, 0x65, 0x6c, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44,
	0x44, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x44, 0x44, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79,
	0x6f, 0x66, 0x66, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x45, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6f, 0x66, 0x66, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x70, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x53, 0x68, 0x61, 0x72,
	0x70, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x69, 0x6e, 0x4d, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x21, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0e, 0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x26, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x22, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x65, 0x61, 0x6c, 0x73,
	0x18, 0x23, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x44, 0x65, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x24, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x25, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x6f, 0x73,
	0x73, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x26, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c,
	0x6f, 0x73, 0x73, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x27, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4c,
	0x6f, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x4c, 0x6f, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x2a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x41, 0x76, 0x67, 0x63, 0x6f, 0x6e, 0x18, 0x2b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x41,
	0x76, 0x67, 0x63, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x4c, 0x6f, 0x73, 0x73, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x41, 0x76, 0x67, 0x43, 0x6f, 0x6e, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x4c, 0x6f, 0x73, 0x73, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x41, 0x76, 0x67, 0x43, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x67, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x2d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4d, 0x61, 0x67, 0x69, 0x63, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x2e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x48, 0x0a,
	0x09, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x2f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x09, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x2d, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_testingstatistics_proto_rawDescOnce sync.Once
	file_testingstatistics_proto_rawDescData = file_testingstatistics_proto_rawDesc
)

func file_testingstatistics_proto_rawDescGZIP() []byte {
	file_testingstatistics_proto_rawDescOnce.Do(func() {
		file_testingstatistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_testingstatistics_proto_rawDescData)
	})
	return file_testingstatistics_proto_rawDescData
}

var file_testingstatistics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testingstatistics_proto_goTypes = []interface{}{
	(*TestingStatisticsVariables)(nil), // 0: tickerbeats.v1.TestingStatisticsVariables
	(*TestingStatistics)(nil),          // 1: tickerbeats.v1.TestingStatistics
}
var file_testingstatistics_proto_depIdxs = []int32{
	0, // 0: tickerbeats.v1.TestingStatistics.Variables:type_name -> tickerbeats.v1.TestingStatisticsVariables
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_testingstatistics_proto_init() }
func file_testingstatistics_proto_init() {
	if File_testingstatistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testingstatistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestingStatisticsVariables); i {
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
		file_testingstatistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestingStatistics); i {
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
			RawDescriptor: file_testingstatistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testingstatistics_proto_goTypes,
		DependencyIndexes: file_testingstatistics_proto_depIdxs,
		MessageInfos:      file_testingstatistics_proto_msgTypes,
	}.Build()
	File_testingstatistics_proto = out.File
	file_testingstatistics_proto_rawDesc = nil
	file_testingstatistics_proto_goTypes = nil
	file_testingstatistics_proto_depIdxs = nil
}
