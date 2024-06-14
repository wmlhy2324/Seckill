package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_stock/stock_rpc/internal/config"
	"github.com/go-redis/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	RedisClient := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		Redis:  RedisClient,
	}
}
