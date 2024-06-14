package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_cart/cart_rpc/cart"
	"Final_Assessment/lhy_order/order_rpc/internal/config"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config               config.Config
	DB                   *gorm.DB
	CartRpc              cart.Cart
	AsynqClient          *asynq.Client
	KqueuePutOrderClient *kq.Pusher
	Redis                *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB := core.InitGorm(c.Mysql.DataSource)
	RedisDB := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:               c,
		DB:                   MysqlDB,
		CartRpc:              cart.NewCart(zrpc.MustNewClient(c.CartRpc)),
		AsynqClient:          newAsynqClient(c),
		KqueuePutOrderClient: kq.NewPusher(c.KqPutOrderConf.Brokers, c.KqPutOrderConf.Topic),
		Redis:                RedisDB,
	}
}
