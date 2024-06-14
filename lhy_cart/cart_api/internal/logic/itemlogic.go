package logic

import (
	"Final_Assessment/lhy_cart/cart_api/internal/svc"
	"Final_Assessment/lhy_cart/cart_api/internal/types"
	"Final_Assessment/lhy_cart/cart_models"
	"Final_Assessment/utils/jwts"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"strconv"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var rng = rand.New(src)
var minPExpire = 1 * time.Minute
var maxPExpire = 10 * time.Minute
var randomDuration = time.Duration(rng.Int63n(int64(maxPExpire.Minutes()-minPExpire.Minutes()+1))+int64(minPExpire.Minutes())) * time.Minute

type ItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ItemLogic {
	return &ItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ItemLogic) Item(req *types.CreateItemRequest) (resp *types.CreateItemResponse, err error) {
	payload, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	//这里2设置的是普通用户权限，如果是管理员权限可以设置为1
	if payload.Role != 1 {
		err = errors.New("权限不足")
		return
	}
	var item cart_models.ItemModel
	var SeckillItem cart_models.SeckillItemModel
	err = l.svcCtx.DB.Take(&item, "item_name = ?", req.ItemName).Error
	if err == nil {
		err = errors.New("商品已存在")
		return
	}

	if req.Type == 0 {

		item = cart_models.ItemModel{
			Name:        req.ItemName,
			Price:       req.Price,
			Stock:       uint(req.Stock),
			Description: req.Description,
		}
		err = l.svcCtx.DB.Create(&item).Error
		if err != nil {
			err = errors.New("商品创建失败")
			return

		}
		return &types.CreateItemResponse{
			ItemName:    item.Name,
			Description: item.Description,
			Price:       item.Price,
		}, nil
	}
	if req.Type == 1 {
		SeckillItem = cart_models.SeckillItemModel{
			Name:        req.ItemName,
			Price:       req.Price,
			Stock:       uint(req.Stock),
			Description: req.Description,
		}

		err = l.svcCtx.DB.Create(&SeckillItem).Error
		if err != nil {
			err = errors.New("商品创建失败")
			return

		}
		cacheKey := "itemId:" + strconv.Itoa(int(SeckillItem.ID))
		_, err = l.svcCtx.Redis.HMSet(cacheKey, map[string]interface{}{
			"name":        SeckillItem.Name,
			"price":       SeckillItem.Price,
			"stock":       SeckillItem.Stock,
			"description": SeckillItem.Description,
		}).Result()
		if err != nil {
			return nil, errors.New("写入缓存失败")
		}
		err = l.svcCtx.Redis.PExpire(cacheKey, randomDuration).Err()
		if err != nil {
			return nil, errors.New("设置过期时间失败")
		}

		return &types.CreateItemResponse{
			ItemName:    SeckillItem.Name,
			Description: SeckillItem.Description,
			Price:       SeckillItem.Price,
		}, nil
	}

	return nil, errors.New("商品类型不正确")
}
