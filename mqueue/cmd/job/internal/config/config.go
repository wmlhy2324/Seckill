package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	CartRpc  zrpc.RpcClientConf
	OrderRpc zrpc.RpcClientConf
	Redis    struct {
		Addr     string
		Password string
		DB       int
		Type     string
	}
}
