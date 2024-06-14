package logic

import (
	"context"
	"fmt"
	"strconv"

	"Final_Assessment/lhy_stock/stock_rpc/internal/svc"
	"Final_Assessment/lhy_stock/stock_rpc/types/stock_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReduceInventoryRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReduceInventoryRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReduceInventoryRollBackLogic {
	return &ReduceInventoryRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReduceInventoryRollBackLogic) ReduceInventoryRollBack(in *stock_rpc.ReduceInventoryRequest) (*stock_rpc.ReduceInventoryResponse, error) {
	fmt.Println("减库存回滚")
	key := prefix + strconv.Itoa(int(in.ItemId))
	_, err := l.svcCtx.Redis.Incr(key).Result()
	if err != nil {
		logx.Info("incr key %s failed: %s", key, err)
		return nil, err
	}

	return &stock_rpc.ReduceInventoryResponse{}, nil
}
