package logic

import (
	"Final_Assessment/lhy_user/user_models"
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"Final_Assessment/lhy_user/user_rpc/internal/svc"
	"Final_Assessment/lhy_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user_rpc.UserInfoRequest) (*user_rpc.UserInfoResponse, error) {
	var user user_models.UserModel
	// todo: add your logic here and delete this line
	cacheKey := "userId:" + strconv.Itoa(int(in.UserId))
	dataUser, err := l.svcCtx.Redis.HGetAll(cacheKey).Result()
	if err != nil {
		return nil, errors.New("缓存获取失败")
	}
	if len(dataUser) == 0 {
		l.svcCtx.DB.Where("id = ?", in.UserId).Find(user)
		if user.ID == 0 {
			return nil, errors.New("用户没有找到")
		}
		l.svcCtx.Redis.HMSet(cacheKey, map[string]interface{}{
			"Account": user.Account,
			"Role":    user.Role,
		})
		data, _ := json.Marshal(user)
		return &user_rpc.UserInfoResponse{Data: data}, nil
	}
	data, _ := json.Marshal(dataUser)
	return &user_rpc.UserInfoResponse{Data: data}, nil
}
