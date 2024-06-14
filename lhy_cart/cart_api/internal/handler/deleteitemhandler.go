package handler

import (
	"Final_Assessment/common/response"
	"Final_Assessment/lhy_cart/cart_api/internal/logic"
	"Final_Assessment/lhy_cart/cart_api/internal/svc"
	"Final_Assessment/lhy_cart/cart_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteItemRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewDeleteItemLogic(r.Context(), svcCtx)
		resp, err := l.DeleteItem(&req)
		response.Response(r, w, resp, err)

	}
}
