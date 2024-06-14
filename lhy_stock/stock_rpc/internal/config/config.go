package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}
