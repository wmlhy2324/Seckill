package handler

import (
	"Final_Assessment/common/response"
	"Final_Assessment/lhy_cart/cart_api/internal/logic"
	"Final_Assessment/lhy_cart/cart_api/internal/svc"
	"Final_Assessment/lhy_cart/cart_models"
	"encoding/json"
	"errors"
	"net/http"
)

func getItemsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetItemsLogic(r.Context(), svcCtx)
		var item []cart_models.ItemModel
		token := r.Header.Get("token")
		resp, err := l.GetItems(token)
		err = json.Unmarshal(resp.Data, &item)
		if err != nil {
			err = errors.New("获取商品失败")
		}
		response.Response(r, w, item, err)

	}
}
