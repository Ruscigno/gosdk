// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tickerbeats_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TickerBeatsServiceClient is the client API for TickerBeatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TickerBeatsServiceClient interface {
	GetTickerBeats(ctx context.Context, in *TickerBeatsRequest, opts ...grpc.CallOption) (*TradeBeatsResponse, error)
}

type tickerBeatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTickerBeatsServiceClient(cc grpc.ClientConnInterface) TickerBeatsServiceClient {
	return &tickerBeatsServiceClient{cc}
}

func (c *tickerBeatsServiceClient) GetTickerBeats(ctx context.Context, in *TickerBeatsRequest, opts ...grpc.CallOption) (*TradeBeatsResponse, error) {
	out := new(TradeBeatsResponse)
	err := c.cc.Invoke(ctx, "/tickerbeats.v1.TickerBeatsService/GetTickerBeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TickerBeatsServiceServer is the server API for TickerBeatsService service.
// All implementations must embed UnimplementedTickerBeatsServiceServer
// for forward compatibility
type TickerBeatsServiceServer interface {
	GetTickerBeats(context.Context, *TickerBeatsRequest) (*TradeBeatsResponse, error)
	mustEmbedUnimplementedTickerBeatsServiceServer()
}

// UnimplementedTickerBeatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTickerBeatsServiceServer struct {
}

func (UnimplementedTickerBeatsServiceServer) GetTickerBeats(context.Context, *TickerBeatsRequest) (*TradeBeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickerBeats not implemented")
}
func (UnimplementedTickerBeatsServiceServer) mustEmbedUnimplementedTickerBeatsServiceServer() {}

// UnsafeTickerBeatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TickerBeatsServiceServer will
// result in compilation errors.
type UnsafeTickerBeatsServiceServer interface {
	mustEmbedUnimplementedTickerBeatsServiceServer()
}

func RegisterTickerBeatsServiceServer(s grpc.ServiceRegistrar, srv TickerBeatsServiceServer) {
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
		FullMethod: "/tickerbeats.v1.TickerBeatsService/GetTickerBeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickerBeatsServiceServer).GetTickerBeats(ctx, req.(*TickerBeatsRequest))
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// TransactionsServiceClient is the client API for TransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
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
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tickerbeats.v1.TransactionsService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateDeals(ctx context.Context, in *CreateDealsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tickerbeats.v1.TransactionsService/CreateDeals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tickerbeats.v1.TransactionsService/CreateOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreatePositions(ctx context.Context, in *CreatePositionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tickerbeats.v1.TransactionsService/CreatePositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateTradeTransaction(ctx context.Context, in *CreateTradeTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tickerbeats.v1.TransactionsService/CreateTradeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServiceServer is the server API for TransactionsService service.
// All implementations must embed UnimplementedTransactionsServiceServer
// for forward compatibility
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

// UnimplementedTransactionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionsServiceServer struct {
}

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

// UnsafeTransactionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionsServiceServer will
// result in compilation errors.
type UnsafeTransactionsServiceServer interface {
	mustEmbedUnimplementedTransactionsServiceServer()
}

func RegisterTransactionsServiceServer(s grpc.ServiceRegistrar, srv TransactionsServiceServer) {
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
		FullMethod: "/tickerbeats.v1.TransactionsService/CreateAccount",
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
		FullMethod: "/tickerbeats.v1.TransactionsService/CreateDeals",
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
		FullMethod: "/tickerbeats.v1.TransactionsService/CreateOrders",
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
		FullMethod: "/tickerbeats.v1.TransactionsService/CreatePositions",
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
		FullMethod: "/tickerbeats.v1.TransactionsService/CreateTradeTransaction",
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
