package svc

import (
	"Final_Assessment/lhy_cart/cart_rpc/cart"
	"Final_Assessment/lhy_order/order_rpc/order"
	"Final_Assessment/mqueue/cmd/job/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	OrderRpc    order.Order
	CartRpc     cart.Cart
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		OrderRpc:    order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		CartRpc:     cart.NewCart(zrpc.MustNewClient(c.CartRpc)),
	}
}
