// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"userservice/userclient"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.Healthreq) (resp *types.Healthres, err error) {
	// todo: add your logic here and delete this line
	var resprpc *userclient.Response
	if req.Id == 1 {
		resprpc, err = l.svcCtx.UserRpc.Ping(l.ctx, &userclient.Request{Ping: "ping"})
	} else {
		resprpc = &userclient.Response{
			Pong: "beng",
		}
	}
	resp = &types.Healthres{
		Msg: resprpc.Pong,
	}
	return
}
