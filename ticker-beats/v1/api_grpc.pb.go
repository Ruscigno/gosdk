// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TickerBeatsService_GetTickerBeats_FullMethodName             = "/tickerbeats.v1.TickerBeatsService/GetTickerBeats"
	TickerBeatsService_GetTickerBeatsNotification_FullMethodName = "/tickerbeats.v1.TickerBeatsService/GetTickerBeatsNotification"
	TickerBeatsService_ConfirmBeatsTransaction_FullMethodName    = "/tickerbeats.v1.TickerBeatsService/ConfirmBeatsTransaction"
)

// TickerBeatsServiceClient is the client API for TickerBeatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for TickerBeatsService Microservice
type TickerBeatsServiceClient interface {
	GetTickerBeats(ctx context.Context, in *TickerBeatsRequest, opts ...grpc.CallOption) (*TickerBeatsResponse, error)
	GetTickerBeatsNotification(ctx context.Context, in *TickerBeatsRequest, opts ...grpc.CallOption) (*TickerBeatsResponse, error)
	ConfirmBeatsTransaction(ctx context.Context, in *TickerBeatsConfirmationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tickerBeatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTickerBeatsServiceClient(cc grpc.ClientConnInterface) TickerBeatsServiceClient {
	return &tickerBeatsServiceClient{cc}
}

func (c *tickerBeatsServiceClient) GetTickerBeats(ctx context.Context, in *TickerBeatsRequest, opts ...grpc.CallOption) (*TickerBeatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TickerBeatsResponse)
	err := c.cc.Invoke(ctx, TickerBeatsService_GetTickerBeats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tickerBeatsServiceClient) GetTickerBeatsNotification(ctx context.Context, in *TickerBeatsRequest, opts ...grpc.CallOption) (*TickerBeatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TickerBeatsResponse)
	err := c.cc.Invoke(ctx, TickerBeatsService_GetTickerBeatsNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tickerBeatsServiceClient) ConfirmBeatsTransaction(ctx context.Context, in *TickerBeatsConfirmationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TickerBeatsService_ConfirmBeatsTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TickerBeatsServiceServer is the server API for TickerBeatsService service.
// All implementations must embed UnimplementedTickerBeatsServiceServer
// for forward compatibility.
//
// Service for TickerBeatsService Microservice
type TickerBeatsServiceServer interface {
	GetTickerBeats(context.Context, *TickerBeatsRequest) (*TickerBeatsResponse, error)
	GetTickerBeatsNotification(context.Context, *TickerBeatsRequest) (*TickerBeatsResponse, error)
	ConfirmBeatsTransaction(context.Context, *TickerBeatsConfirmationRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTickerBeatsServiceServer()
}

// UnimplementedTickerBeatsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTickerBeatsServiceServer struct{}

func (UnimplementedTickerBeatsServiceServer) GetTickerBeats(context.Context, *TickerBeatsRequest) (*TickerBeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickerBeats not implemented")
}
func (UnimplementedTickerBeatsServiceServer) GetTickerBeatsNotification(context.Context, *TickerBeatsRequest) (*TickerBeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickerBeatsNotification not implemented")
}
func (UnimplementedTickerBeatsServiceServer) ConfirmBeatsTransaction(context.Context, *TickerBeatsConfirmationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmBeatsTransaction not implemented")
}
func (UnimplementedTickerBeatsServiceServer) mustEmbedUnimplementedTickerBeatsServiceServer() {}
func (UnimplementedTickerBeatsServiceServer) testEmbeddedByValue()                            {}

// UnsafeTickerBeatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TickerBeatsServiceServer will
// result in compilation errors.
type UnsafeTickerBeatsServiceServer interface {
	mustEmbedUnimplementedTickerBeatsServiceServer()
}

func RegisterTickerBeatsServiceServer(s grpc.ServiceRegistrar, srv TickerBeatsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTickerBeatsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TickerBeatsService_ServiceDesc, srv)
}

func _TickerBeatsService_GetTickerBeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickerBeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickerBeatsServiceServer).GetTickerBeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickerBeatsService_GetTickerBeats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickerBeatsServiceServer).GetTickerBeats(ctx, req.(*TickerBeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TickerBeatsService_GetTickerBeatsNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickerBeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickerBeatsServiceServer).GetTickerBeatsNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickerBeatsService_GetTickerBeatsNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickerBeatsServiceServer).GetTickerBeatsNotification(ctx, req.(*TickerBeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TickerBeatsService_ConfirmBeatsTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickerBeatsConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickerBeatsServiceServer).ConfirmBeatsTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickerBeatsService_ConfirmBeatsTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickerBeatsServiceServer).ConfirmBeatsTransaction(ctx, req.(*TickerBeatsConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TickerBeatsService_ServiceDesc is the grpc.ServiceDesc for TickerBeatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TickerBeatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tickerbeats.v1.TickerBeatsService",
	HandlerType: (*TickerBeatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTickerBeats",
			Handler:    _TickerBeatsService_GetTickerBeats_Handler,
		},
		{
			MethodName: "GetTickerBeatsNotification",
			Handler:    _TickerBeatsService_GetTickerBeatsNotification_Handler,
		},
		{
			MethodName: "ConfirmBeatsTransaction",
			Handler:    _TickerBeatsService_ConfirmBeatsTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	TransactionsService_CreateAccount_FullMethodName          = "/tickerbeats.v1.TransactionsService/CreateAccount"
	TransactionsService_CreateDeals_FullMethodName            = "/tickerbeats.v1.TransactionsService/CreateDeals"
	TransactionsService_CreateOrders_FullMethodName           = "/tickerbeats.v1.TransactionsService/CreateOrders"
	TransactionsService_CreatePositions_FullMethodName        = "/tickerbeats.v1.TransactionsService/CreatePositions"
	TransactionsService_CreateTradeTransaction_FullMethodName = "/tickerbeats.v1.TransactionsService/CreateTradeTransaction"
)

// TransactionsServiceClient is the client API for TransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for TickerBeatsService Microservice
type TransactionsServiceClient interface {
	// Creates a new account
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Creates a new deal
	CreateDeals(ctx context.Context, in *CreateDealsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Creates a new order
	CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Creates a new position
	CreatePositions(ctx context.Context, in *CreatePositionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Creates a trade transation
	CreateTradeTransaction(ctx context.Context, in *CreateTradeTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type transactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionsServiceClient(cc grpc.ClientConnInterface) TransactionsServiceClient {
	return &transactionsServiceClient{cc}
}

func (c *transactionsServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TransactionsService_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateDeals(ctx context.Context, in *CreateDealsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TransactionsService_CreateDeals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TransactionsService_CreateOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreatePositions(ctx context.Context, in *CreatePositionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TransactionsService_CreatePositions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateTradeTransaction(ctx context.Context, in *CreateTradeTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TransactionsService_CreateTradeTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServiceServer is the server API for TransactionsService service.
// All implementations must embed UnimplementedTransactionsServiceServer
// for forward compatibility.
//
// Service for TickerBeatsService Microservice
type TransactionsServiceServer interface {
	// Creates a new account
	CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error)
	// Creates a new deal
	CreateDeals(context.Context, *CreateDealsRequest) (*emptypb.Empty, error)
	// Creates a new order
	CreateOrders(context.Context, *CreateOrdersRequest) (*emptypb.Empty, error)
	// Creates a new position
	CreatePositions(context.Context, *CreatePositionsRequest) (*emptypb.Empty, error)
	// Creates a trade transation
	CreateTradeTransaction(context.Context, *CreateTradeTransactionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTransactionsServiceServer()
}

// UnimplementedTransactionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionsServiceServer struct{}

func (UnimplementedTransactionsServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedTransactionsServiceServer) CreateDeals(context.Context, *CreateDealsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeals not implemented")
}
func (UnimplementedTransactionsServiceServer) CreateOrders(context.Context, *CreateOrdersRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedTransactionsServiceServer) CreatePositions(context.Context, *CreatePositionsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePositions not implemented")
}
func (UnimplementedTransactionsServiceServer) CreateTradeTransaction(context.Context, *CreateTradeTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTradeTransaction not implemented")
}
func (UnimplementedTransactionsServiceServer) mustEmbedUnimplementedTransactionsServiceServer() {}
func (UnimplementedTransactionsServiceServer) testEmbeddedByValue()                             {}

// UnsafeTransactionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionsServiceServer will
// result in compilation errors.
type UnsafeTransactionsServiceServer interface {
	mustEmbedUnimplementedTransactionsServiceServer()
}

func RegisterTransactionsServiceServer(s grpc.ServiceRegistrar, srv TransactionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionsService_ServiceDesc, srv)
}

func _TransactionsService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionsService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_CreateDeals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDealsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreateDeals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionsService_CreateDeals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreateDeals(ctx, req.(*CreateDealsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_CreateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionsService_CreateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreateOrders(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_CreatePositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreatePositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionsService_CreatePositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreatePositions(ctx, req.(*CreatePositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_CreateTradeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTradeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreateTradeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionsService_CreateTradeTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreateTradeTransaction(ctx, req.(*CreateTradeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionsService_ServiceDesc is the grpc.ServiceDesc for TransactionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tickerbeats.v1.TransactionsService",
	HandlerType: (*TransactionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _TransactionsService_CreateAccount_Handler,
		},
		{
			MethodName: "CreateDeals",
			Handler:    _TransactionsService_CreateDeals_Handler,
		},
		{
			MethodName: "CreateOrders",
			Handler:    _TransactionsService_CreateOrders_Handler,
		},
		{
			MethodName: "CreatePositions",
			Handler:    _TransactionsService_CreatePositions_Handler,
		},
		{
			MethodName: "CreateTradeTransaction",
			Handler:    _TransactionsService_CreateTradeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	TestingStatisticsService_Create_FullMethodName = "/tickerbeats.v1.TestingStatisticsService/Create"
)

// TestingStatisticsServiceClient is the client API for TestingStatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for TestingStatistics Microservice
type TestingStatisticsServiceClient interface {
	Create(ctx context.Context, in *CreateTestingStatisticsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type testingStatisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestingStatisticsServiceClient(cc grpc.ClientConnInterface) TestingStatisticsServiceClient {
	return &testingStatisticsServiceClient{cc}
}

func (c *testingStatisticsServiceClient) Create(ctx context.Context, in *CreateTestingStatisticsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TestingStatisticsService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestingStatisticsServiceServer is the server API for TestingStatisticsService service.
// All implementations must embed UnimplementedTestingStatisticsServiceServer
// for forward compatibility.
//
// Service for TestingStatistics Microservice
type TestingStatisticsServiceServer interface {
	Create(context.Context, *CreateTestingStatisticsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTestingStatisticsServiceServer()
}

// UnimplementedTestingStatisticsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestingStatisticsServiceServer struct{}

func (UnimplementedTestingStatisticsServiceServer) Create(context.Context, *CreateTestingStatisticsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTestingStatisticsServiceServer) mustEmbedUnimplementedTestingStatisticsServiceServer() {
}
func (UnimplementedTestingStatisticsServiceServer) testEmbeddedByValue() {}

// UnsafeTestingStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestingStatisticsServiceServer will
// result in compilation errors.
type UnsafeTestingStatisticsServiceServer interface {
	mustEmbedUnimplementedTestingStatisticsServiceServer()
}

func RegisterTestingStatisticsServiceServer(s grpc.ServiceRegistrar, srv TestingStatisticsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestingStatisticsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestingStatisticsService_ServiceDesc, srv)
}

func _TestingStatisticsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestingStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingStatisticsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingStatisticsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingStatisticsServiceServer).Create(ctx, req.(*CreateTestingStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestingStatisticsService_ServiceDesc is the grpc.ServiceDesc for TestingStatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestingStatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tickerbeats.v1.TestingStatisticsService",
	HandlerType: (*TestingStatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TestingStatisticsService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
