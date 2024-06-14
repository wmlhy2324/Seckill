package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"

	"Final_Assessment/lhy_stock/stock_rpc/internal/svc"
	"Final_Assessment/lhy_stock/stock_rpc/types/stock_rpc"

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

func (l *ReduceInventoryLogic) ReduceInventory(in *stock_rpc.ReduceInventoryRequest) (*stock_rpc.ReduceInventoryResponse, error) {
	fmt.Println("开始减库存")
	key := prefix + strconv.Itoa(int(in.ItemId))
	n, err := l.svcCtx.Redis.Decr(key).Result()
	if err != nil {
		logx.Info("decr key %s failed: %s", key, err)
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		if n < 0 {
			logx.Info("%d已无库存，减1失败", in.ItemId)
			return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
		}

	}
	return nil, nil

}
