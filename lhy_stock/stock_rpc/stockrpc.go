package main

import (
	"flag"
	"fmt"

	"Final_Assessment/lhy_stock/stock_rpc/internal/config"
	"Final_Assessment/lhy_stock/stock_rpc/internal/server"
	"Final_Assessment/lhy_stock/stock_rpc/internal/svc"
	"Final_Assessment/lhy_stock/stock_rpc/types/stock_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/stockrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		stock_rpc.RegisterStockServer(grpcServer, server.NewStockServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
