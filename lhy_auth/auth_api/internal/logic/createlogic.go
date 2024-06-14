package logic

import (
	"Final_Assessment/lhy_auth/auth_api/internal/svc"
	"Final_Assessment/lhy_auth/auth_api/internal/types"
	"Final_Assessment/lhy_auth/auth_models"
	"Final_Assessment/lhy_user/user_rpc/types/user_rpc"
	"Final_Assessment/utils/pwd"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", req.Nickname).Error
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	hashPwd := pwd.HashPwd(req.Password)
	res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
		Username:       req.Nickname,
		Password:       hashPwd,
		Role:           2,
		RegisterSource: "local",
		Account:        500,
	})
	user.ID = uint(res.UserId)
	user.RegisterSource = "local"
	user.Nickname = req.Nickname
	user.Pwd = req.Password

	return &types.CreateResponse{
		ID:       int64(user.ID),
		Nickname: user.Nickname,
		Role:     2,
	}, nil
}
