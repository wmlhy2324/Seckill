package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CartRpc zrpc.RpcClientConf
	Redis   struct {
		Addr     string
		Password string
		DB       int
	}
	KqPutOrderConf KqConfig
}
