package logic

import (
	"context"

	"Final_Assessment/lhy_cart/cart_api/internal/svc"
	"Final_Assessment/lhy_cart/cart_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemLogic {
	return &DeleteItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteItemLogic) DeleteItem(req *types.DeleteItemRequest) (resp *types.DeleteItemResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
