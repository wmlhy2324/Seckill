package kq

import (
	"Final_Assessment/lhy_order/mq/internal/svc"
	"Final_Assessment/lhy_order/order_models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type PutOrder struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutOrder(ctx context.Context, svcCtx *svc.ServiceContext) *PutOrder {
	return &PutOrder{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *PutOrder) Consume(_, val string) error {
	logx.Infof("PaymentSuccess key : val :%s", val)
	var order order_models.OrderModel
	err := json.Unmarshal([]byte(val), &order)
	if err != nil {
		logx.WithContext(l.ctx).Error("consume order fail : %v , val : %s", err, val)
	}
	err = l.svcCtx.DB.Create(&order).Error
	if err != nil {
		logx.WithContext(l.ctx).Error("create order fail : %v ", err)
		return err
	}
	cacheKey := "OrderId:" + strconv.Itoa(int(order.ID))
	l.svcCtx.Redis.HMSet(cacheKey, map[string]interface{}{
		"userId": order.UserId,
		"amount": order.Amount,
	})
	l.svcCtx.Redis.PExpire(cacheKey, time.Minute*1)

	return nil
}
