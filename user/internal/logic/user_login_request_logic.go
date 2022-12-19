package logic

import (
	"context"
	"errors"
	"myiot/helper"
	"myiot/models"

	"myiot/user/internal/svc"
	"myiot/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginRequestLogic {
	return &UserLoginRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginRequestLogic) UserLoginRequest(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserLoginReply)
	ub := new(models.UserBasic)
	err = l.svcCtx.DB.Where("username = ? AND password = ?", req.Username, helper.Md5(req.Password)).First(ub).Error
	if err != nil {
		logx.Error("[DB ERROR :", err)
		err = errors.New("username or password has error")
		return
	}
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Name, 3600*24*30)
	if err != nil {
		logx.Error("[DB ERROR :", err)
		err = errors.New("username or password has error")
		return
	}
	resp.Token = token
	return
}
