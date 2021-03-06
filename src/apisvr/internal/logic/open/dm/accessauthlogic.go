package dm

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/dm"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessAuthLogic {
	return &AccessAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccessAuthLogic) AccessAuth(req *types.AccessAuthReq) error {
	l.Infof("AccessAuth|req=%+v", req)
	access := req.Access
	//如果是
	switch req.Access {
	case "1":
		access = def.SUB
	case "2":
		access = def.PUB
	}
	_, err := l.svcCtx.DmRpc.AccessAuth(l.ctx, &dm.AccessAuthReq{
		Username: req.Username,
		Topic:    req.Topic,
		ClientID: req.ClientID,
		Access:   access,
		Ip:       req.Ip,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s|rpc.AccessAuth|req=%v|err=%+v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
