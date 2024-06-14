package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd string
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	CartRpc  zrpc.RpcClientConf
	OrderRpc zrpc.RpcClientConf
}
