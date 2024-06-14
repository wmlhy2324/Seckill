// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"Final_Assessment/lhy_order/order_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/item",
				Handler: CreateOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/seckillitem",
				Handler: CreateSeckillOrderHandler(serverCtx),
			},
		},
	)
}
