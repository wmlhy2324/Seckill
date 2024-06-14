package logic

import (
	"context"
	"strconv"

	"Final_Assessment/lhy_cart/cart_rpc/internal/svc"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseInventoryLogic {
	return &IncreaseInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseInventoryLogic) IncreaseInventory(in *cart_rpc.IncreaseInventoryRequest) (*cart_rpc.IncreaseInventoryResponse, error) {

	key := prefix + strconv.Itoa(int(in.ItemId))
	_, err := l.svcCtx.Redis.Incr(key).Result()
	if err != nil {
		logx.Info("incr key %s failed: %s", key, err)
		return nil, err
	}

	return &cart_rpc.IncreaseInventoryResponse{}, nil
}
