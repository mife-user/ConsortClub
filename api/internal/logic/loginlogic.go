// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// todo: add your logic here and delete this line
	resp = &types.LoginRes{}
	if req.UserName == "" {
		err = errors.New("缺少用户名称")
		return
	} else if len(req.UserName) > 16 {
		err = errors.New("用户名称过长")
		return
	}
	if req.PassWord == "" {
		err = errors.New("缺少密码")
		return
	} else if len(req.UserName) > 32 {
		err = errors.New("密码过长")
		return
	}
	resp.Role = "yyy"
	return
}
