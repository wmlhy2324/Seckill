package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_cart/cart_api/internal/config"
	"Final_Assessment/lhy_cart/cart_rpc/cart"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"
	"Final_Assessment/lhy_order/order_rpc/order"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	Redis    *redis.Client
	CartRpc  cart_rpc.CartClient
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:   c,
		DB:       MysqlDB,
		Redis:    client,
		CartRpc:  cart.NewCart(zrpc.MustNewClient(c.CartRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
