package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Etcd string
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}

	UserRpc   zrpc.RpcClientConf
	WriteList []string
}
