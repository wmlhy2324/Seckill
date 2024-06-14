package logic

import (
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_user/user_models"
	"context"
	"errors"
	"strconv"

	"Final_Assessment/lhy_payment/payment_rpc/internal/svc"
	"Final_Assessment/lhy_payment/payment_rpc/types/payment_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductAccountLogic {
	return &DeductAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductAccountLogic) DeductAccount(in *payment_rpc.DeductAccountRequest) (*payment_rpc.DeductAccountResponse, error) {
	user := new(user_models.UserModel)
	price, err := l.Price(in.OrderId, in.UserId)
	if err != nil {
		return nil, err
	}
	if price == 0 {
		return nil, errors.New("价格有误")
	}

	user.Account -= price
	if user.Account < 0 {
		return nil, errors.New("余额不足")
	}
	err = l.svcCtx.DB.Model(&user_models.UserModel{}).Where("user_id = ?", in.UserId).Update("amount", user.Account).Error
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DB.Model(&order_models.OrderModel{}).Where("user_id = ?", in.UserId).Update("status", 1).Error
	if err != nil {
		return nil, err
	}
	cacheKey := "userId:" + strconv.Itoa(int(in.UserId))
	l.svcCtx.Redis.HSet(cacheKey, "amount", user.Account)
	return &payment_rpc.DeductAccountResponse{
		Account: float32(user.Account),
	}, nil
}
func (l *DeductAccountLogic) Price(OrderId, UserId int32) (float64, error) {
	var order order_models.OrderModel
	cacheKey := "OrderId:" + strconv.Itoa(int(OrderId))
	DataOrder, err := l.svcCtx.Redis.HGetAll(cacheKey).Result()
	if err != nil {
		return 0, errors.New("获取订单缓存失败")
	}
	if len(DataOrder) == 0 {
		l.svcCtx.DB.Where("id = ? AND userid = ? AND status = ?", OrderId, UserId, 0).Find(&order)
		if order.ID == 0 {
			return 0, errors.New("订单没有找到")
		}
	}
	priceStr, ok := DataOrder["amount"]
	if !ok {
		// 价格字段不存在
		return 0, errors.New("字段错误")
	}

	// 将字符串转换为浮点数
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, errors.New("转换类型错误")
	}
	return price, nil
}
