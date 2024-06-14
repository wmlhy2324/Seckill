// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: stock_rpc.proto

package stock_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Stock_ReduceInventory_FullMethodName         = "/stock_rpc.stock/ReduceInventory"
	Stock_ReduceInventoryRollBack_FullMethodName = "/stock_rpc.stock/ReduceInventoryRollBack"
)

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockClient interface {
	ReduceInventory(ctx context.Context, in *ReduceInventoryRequest, opts ...grpc.CallOption) (*ReduceInventoryResponse, error)
	ReduceInventoryRollBack(ctx context.Context, in *ReduceInventoryRequest, opts ...grpc.CallOption) (*ReduceInventoryResponse, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) ReduceInventory(ctx context.Context, in *ReduceInventoryRequest, opts ...grpc.CallOption) (*ReduceInventoryResponse, error) {
	out := new(ReduceInventoryResponse)
	err := c.cc.Invoke(ctx, Stock_ReduceInventory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) ReduceInventoryRollBack(ctx context.Context, in *ReduceInventoryRequest, opts ...grpc.CallOption) (*ReduceInventoryResponse, error) {
	out := new(ReduceInventoryResponse)
	err := c.cc.Invoke(ctx, Stock_ReduceInventoryRollBack_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
// All implementations must embed UnimplementedStockServer
// for forward compatibility
type StockServer interface {
	ReduceInventory(context.Context, *ReduceInventoryRequest) (*ReduceInventoryResponse, error)
	ReduceInventoryRollBack(context.Context, *ReduceInventoryRequest) (*ReduceInventoryResponse, error)
	mustEmbedUnimplementedStockServer()
}

// UnimplementedStockServer must be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (UnimplementedStockServer) ReduceInventory(context.Context, *ReduceInventoryRequest) (*ReduceInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReduceInventory not implemented")
}
func (UnimplementedStockServer) ReduceInventoryRollBack(context.Context, *ReduceInventoryRequest) (*ReduceInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReduceInventoryRollBack not implemented")
}
func (UnimplementedStockServer) mustEmbedUnimplementedStockServer() {}

// UnsafeStockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServer will
// result in compilation errors.
type UnsafeStockServer interface {
	mustEmbedUnimplementedStockServer()
}

func RegisterStockServer(s grpc.ServiceRegistrar, srv StockServer) {
	s.RegisterService(&Stock_ServiceDesc, srv)
}

func _Stock_ReduceInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReduceInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).ReduceInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stock_ReduceInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).ReduceInventory(ctx, req.(*ReduceInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_ReduceInventoryRollBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReduceInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).ReduceInventoryRollBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stock_ReduceInventoryRollBack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).ReduceInventoryRollBack(ctx, req.(*ReduceInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Stock_ServiceDesc is the grpc.ServiceDesc for Stock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock_rpc.stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReduceInventory",
			Handler:    _Stock_ReduceInventory_Handler,
		},
		{
			MethodName: "ReduceInventoryRollBack",
			Handler:    _Stock_ReduceInventoryRollBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock_rpc.proto",
}
