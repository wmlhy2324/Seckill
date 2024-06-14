package jobtype

import "Final_Assessment/lhy_order/order_models"

type DeferCloseOrderPayload struct {
	UserId int32
	ItemId int32
}
type PaySuccessNotifyUserPayload struct {
	Order *order_models.OrderModel
}
