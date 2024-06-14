package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_user/user_rpc/internal/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB := core.InitGorm(c.Mysql.DataSource)
	RedisClient := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		DB:     MysqlDB,
		Redis:  RedisClient,
	}
}
