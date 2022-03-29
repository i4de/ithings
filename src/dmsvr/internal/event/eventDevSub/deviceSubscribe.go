package eventDevSub

//设备的发布,连接及断连处理
import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceDetail"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceSend"
	"github.com/i-Things/things/src/dmsvr/internal/repo/mysql"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type DeviceMsgHandle struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
	logx.Logger
}

func NewDeviceMsgHandle(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceMsgHandle {
	return &DeviceMsgHandle{
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}

func (l *DeviceMsgHandle) Publish(msg *deviceSend.Elements) error {
	l.Infof("DevReqLogic|req=%+v", msg)
	return NewPublishLogic(l.ctx, l.svcCtx).Handle(msg)
}

func (l *DeviceMsgHandle) Connected(msg *deviceSend.Elements) error {
	l.Infof("ConnectLogic|req=%+v", msg)
	//todo 这里需要查询下数据库,避免数据错误
	ld, err := deviceDetail.GetClientIDInfo(msg.ClientID)
	if err != nil {
		return err
	}
	_, _ = l.svcCtx.DeviceLog.Insert(&mysql.DeviceLog{
		ProductID:   ld.ProductID,
		Action:      msg.Action,
		Timestamp:   time.UnixMilli(msg.Timestamp), // 操作时间
		DeviceName:  ld.DeviceName,
		TranceID:    utils.TraceIdFromContext(l.ctx),
		Content:     string(msg.Payload),
		Topic:       msg.Topic,
		ResultType:  errors.Fmt(err).GetCode(),
		CreatedTime: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (l *DeviceMsgHandle) Disconnected(msg *deviceSend.Elements) error {
	l.Infof("DisconnectLogic|req=%+v", msg)
	ld, err := deviceDetail.GetClientIDInfo(msg.ClientID)
	if err != nil {
		return err
	}
	_, _ = l.svcCtx.DeviceLog.Insert(&mysql.DeviceLog{
		ProductID:   ld.ProductID,
		Action:      msg.Action,
		Timestamp:   time.UnixMilli(msg.Timestamp), // 操作时间
		DeviceName:  ld.DeviceName,
		TranceID:    utils.TraceIdFromContext(l.ctx),
		Content:     string(msg.Payload),
		Topic:       msg.Topic,
		ResultType:  errors.Fmt(err).GetCode(),
		CreatedTime: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}