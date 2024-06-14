package svc

import (
	"Final_Assessment/core"
	"Final_Assessment/lhy_auth/auth_api/internal/config"
	"Final_Assessment/lhy_user/user_rpc/types/user_rpc"
	"Final_Assessment/lhy_user/user_rpc/users"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		DB:      MysqlDB,
		Redis:   client,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
