package logic

import (
	"context"
	"errors"
	"myiot/helper"
	"myiot/user/rpc/user"

	"myiot/user/rpc/internal/svc"
	"myiot/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *pb.UserAuthRequest) (*pb.UserAuthReply, error) {
	// todo: add your logic here and delete this line

	if in.Token == "" {
		return nil, errors.New("token is must")
	}
	userClaim, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthReply)
	resp.Identity = userClaim.Identity
	resp.Id = uint64(userClaim.Id)
	resp.Extend = map[string]string{
		"name": userClaim.Name,
	}
	return resp, nil
}
