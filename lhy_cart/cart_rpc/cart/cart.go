// Code generated by goctl. DO NOT EDIT.
// Source: cart_rpc.proto

package cart

import (
	"context"

	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddItemRequest            = cart_rpc.AddItemRequest
	AddItemResponse           = cart_rpc.AddItemResponse
	CartResponse              = cart_rpc.CartResponse
	CreateCartRequest         = cart_rpc.CreateCartRequest
	GetCartDetailRequest      = cart_rpc.GetCartDetailRequest
	GetCartDetailResponse     = cart_rpc.GetCartDetailResponse
	IncreaseInventoryRequest  = cart_rpc.IncreaseInventoryRequest
	IncreaseInventoryResponse = cart_rpc.IncreaseInventoryResponse
	ReduceInventoryRequest    = cart_rpc.ReduceInventoryRequest
	ReduceInventoryResponse   = cart_rpc.ReduceInventoryResponse

	Cart interface {
		CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
		AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
		GetCartDetail(ctx context.Context, in *GetCartDetailRequest, opts ...grpc.CallOption) (*GetCartDetailResponse, error)
		ReduceInventory(ctx context.Context, in *ReduceInventoryRequest, opts ...grpc.CallOption) (*ReduceInventoryResponse, error)
		IncreaseInventory(ctx context.Context, in *IncreaseInventoryRequest, opts ...grpc.CallOption) (*IncreaseInventoryResponse, error)
	}

	defaultCart struct {
		cli zrpc.Client
	}
)

func NewCart(cli zrpc.Client) Cart {
	return &defaultCart{
		cli: cli,
	}
}

func (m *defaultCart) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart_rpc.NewCartClient(m.cli.Conn())
	return client.CreateCart(ctx, in, opts...)
}

func (m *defaultCart) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	client := cart_rpc.NewCartClient(m.cli.Conn())
	return client.AddItem(ctx, in, opts...)
}

func (m *defaultCart) GetCartDetail(ctx context.Context, in *GetCartDetailRequest, opts ...grpc.CallOption) (*GetCartDetailResponse, error) {
	client := cart_rpc.NewCartClient(m.cli.Conn())
	return client.GetCartDetail(ctx, in, opts...)
}

func (m *defaultCart) ReduceInventory(ctx context.Context, in *ReduceInventoryRequest, opts ...grpc.CallOption) (*ReduceInventoryResponse, error) {
	client := cart_rpc.NewCartClient(m.cli.Conn())
	return client.ReduceInventory(ctx, in, opts...)
}

func (m *defaultCart) IncreaseInventory(ctx context.Context, in *IncreaseInventoryRequest, opts ...grpc.CallOption) (*IncreaseInventoryResponse, error) {
	client := cart_rpc.NewCartClient(m.cli.Conn())
	return client.IncreaseInventory(ctx, in, opts...)
}
