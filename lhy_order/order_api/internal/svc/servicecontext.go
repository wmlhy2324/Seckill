package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_order/order_api/internal/config"
	"Final_Assessment/lhy_order/order_rpc/order"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	OrderRpc order.Order
	Redis    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB := core.InitGorm(c.Mysql.Datasource)
	RedisClient := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:   c,
		DB:       MysqlDB,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		Redis:    RedisClient,
	}
}
