package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OrderRpc zrpc.RpcClientConf
	Redis    struct {
		Addr     string
		Password string
		DB       int
	}
	PutOrderConf kq.KqConf
	Mysql        struct {
		Datasource string
	}
}
