package handler

import (
	"Final_Assessment/common/response"
	"Final_Assessment/lhy_auth/auth_api/internal/logic"
	"Final_Assessment/lhy_auth/auth_api/internal/svc"
	"Final_Assessment/lhy_auth/auth_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func createHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		response.Response(r, w, resp, err)

	}
}
