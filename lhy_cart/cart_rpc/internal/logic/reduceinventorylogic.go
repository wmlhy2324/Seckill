package logic

import (
	"context"
	"errors"
	"strconv"

	"Final_Assessment/lhy_cart/cart_rpc/internal/svc"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefix = "item_count_" //所有key设置统一的前缀，方便后续按前缀遍历key
)

type ReduceInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReduceInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReduceInventoryLogic {
	return &ReduceInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReduceInventoryLogic) ReduceInventory(in *cart_rpc.ReduceInventoryRequest) (*cart_rpc.ReduceInventoryResponse, error) {

	key := prefix + strconv.Itoa(int(in.ItemId))
	n, err := l.svcCtx.Redis.Decr(key).Result()
	if err != nil {
		logx.Info("decr key %s failed: %s", key, err)
		return nil, err
	} else {
		if n < 0 {
			logx.Info("%d已无库存，减1失败", in.ItemId)
			return nil, errors.New("已无库存,减库失败")
		}

	}
	return nil, nil
}
