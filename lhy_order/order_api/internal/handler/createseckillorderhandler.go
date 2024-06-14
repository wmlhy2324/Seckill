package handler

import (
	"Final_Assessment/common/response"
	"Final_Assessment/lhy_order/order_api/internal/logic"
	"Final_Assessment/lhy_order/order_api/internal/svc"
	"Final_Assessment/lhy_order/order_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateSeckillOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSeckillRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewCreateSeckillOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateSeckillOrder(&req)
		response.Response(r, w, resp, err)

	}
}
