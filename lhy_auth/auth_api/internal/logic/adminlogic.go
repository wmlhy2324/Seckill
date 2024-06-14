package logic

import (
	"Final_Assessment/lhy_auth/auth_models"
	"Final_Assessment/lhy_user/user_rpc/types/user_rpc"
	"Final_Assessment/utils/pwd"
	"context"
	"errors"

	"Final_Assessment/lhy_auth/auth_api/internal/svc"
	"Final_Assessment/lhy_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.CreateManagerRequest) (resp *types.CreateManagerResponse, err error) {
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", req.Nickname).Error
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	hashPwd := pwd.HashPwd(req.Password)
	res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
		Username:       req.Nickname,
		Password:       hashPwd,
		Role:           1,
		RegisterSource: "local",
	})
	user.ID = uint(res.UserId)
	user.RegisterSource = "local"
	user.Nickname = req.Nickname
	user.Pwd = req.Password

	return &types.CreateManagerResponse{
		ID:       int64(user.ID),
		Nickname: user.Nickname,
		Role:     1,
	}, nil
}
