package handler

import (
	"Final_Assessment/common/response"
	"Final_Assessment/lhy_auth/auth_api/internal/logic"
	"Final_Assessment/lhy_auth/auth_api/internal/svc"
	"Final_Assessment/lhy_auth/auth_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticationRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication(&req)
		response.Response(r, w, resp, err)

	}
}
