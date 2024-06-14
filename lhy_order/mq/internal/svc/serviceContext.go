package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_order/mq/internal/config"
	"Final_Assessment/lhy_order/order_rpc/order"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"time"
)

const flushInterval = 5 * time.Second

type ServiceContext struct {
	Config   config.Config
	OrderRpc order.Order
	DB       *gorm.DB
	Redis    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB := core.InitGorm(c.Mysql.Datasource)
	RedisClient := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:   c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		DB:       MysqlDB,
		Redis:    RedisClient,
	}
}
