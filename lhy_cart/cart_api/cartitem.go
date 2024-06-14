package main

import (
	"Final_Assessment/common/etcd"
	"Final_Assessment/utils/db"
	"flag"
	"fmt"

	"Final_Assessment/lhy_cart/cart_api/internal/config"
	"Final_Assessment/lhy_cart/cart_api/internal/handler"
	"Final_Assessment/lhy_cart/cart_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/cartitem.yaml", "the config file")

func main() {

	flag.Parse()
	db.InitItemInventory()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	etcd.DeliveryAddress(c.Etcd, c.Name+"_api", fmt.Sprintf("%s:%d", c.Host, c.Port))
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
