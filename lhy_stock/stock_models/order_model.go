package stock_models

import "Final_Assessment/common/models"

type OrderModel struct {
	models.Model
	UserId uint    `json:"userId"` // 用户id
	Amount float64 `json:"amount"` // 金额
	Status uint    `json:"status"` // 订单状态 1.未支付 2.已支付
}
