package order_models

import "Final_Assessment/common/models"

type OrderModel struct {
	models.Model
	UserId uint    `json:"userId"` // 用户id
	Amount float64 `json:"amount"` // 金额
	Status int64   `json:"status"` // 订单状态 0 未支付 1 已支付 -1 取消
}
