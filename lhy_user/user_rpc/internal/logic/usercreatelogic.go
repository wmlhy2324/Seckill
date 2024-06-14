package logic

import (
	"Final_Assessment/lhy_user/user_models"
	"Final_Assessment/utils/expiretime"
	"context"
	"errors"
	"strconv"

	"Final_Assessment/lhy_user/user_rpc/internal/svc"
	"Final_Assessment/lhy_user/user_rpc/types/user_rpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, "nickname = ?", in.Username).Error
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	user = user_models.UserModel{
		Nickname:       in.Username,
		Pwd:            in.Password,
		Role:           int8(in.Role),
		RegisterSource: in.RegisterSource,
		Account:        float64(in.Account),
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		return nil, errors.New("用户创建失败")
	}
	l.svcCtx.DB.Where("nickname = ?", in.Username).Find(&user)
	cacheKey := "userId:" + strconv.Itoa(int(user.ID))
	l.svcCtx.Redis.HMSet(cacheKey, map[string]interface{}{
		"Account": in.Account,
		"Role":    in.Role,
	})
	err = l.svcCtx.Redis.PExpire(cacheKey, expiretime.SetExpireTime()).Err()
	if err != nil {
		return nil, errors.New("设置过期时间失败")
	}

	return &user_rpc.UserCreateResponse{UserId: int32(user.ID)}, nil
}
