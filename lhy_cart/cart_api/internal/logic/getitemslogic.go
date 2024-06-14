package logic

import (
	"Final_Assessment/lhy_cart/cart_models"
	"context"
	"encoding/json"
	"errors"

	"Final_Assessment/lhy_cart/cart_api/internal/svc"
	"Final_Assessment/lhy_cart/cart_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemsLogic {
	return &GetItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetItemsLogic) GetItems(token string) (resp *types.GetItemListResponse, err error) {
	var item []cart_models.ItemModel
	err = l.svcCtx.DB.Where("is_deleted = ?", false).Find(&item).Error
	if err != nil {
		err = errors.New("获取商品失败")
		return nil, err
	}
	byteData, _ := json.Marshal(item)
	return &types.GetItemListResponse{
		Data: byteData,
	}, nil
}
