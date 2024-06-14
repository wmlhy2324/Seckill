package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		Datasource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	Etcd     string
	OrderRpc zrpc.RpcClientConf
	StockRpc zrpc.RpcClientConf
}
