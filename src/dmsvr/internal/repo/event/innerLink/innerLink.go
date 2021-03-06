package innerLink

import (
	"context"
	"github.com/i-Things/things/src/dmsvr/internal/config"
	"github.com/i-Things/things/src/dmsvr/internal/domain/device"
	deviceSend "github.com/i-Things/things/src/dmsvr/internal/domain/service/deviceSend"
)

type (
	InnerLink interface {
		PublishToDev(ctx context.Context, topic string, payload []byte) error
		Subscribe(handle Handle) error
		ReqToDeviceSync(ctx context.Context, reqTopic, respTopic string, req *deviceSend.DeviceReq,
			productID, deviceName string) (*deviceSend.DeviceResp, error)
	}
	Handle        func(ctx context.Context) InnerSubEvent
	InnerSubEvent interface {
		Publish(out *device.PublishMsg) error
		Connected(out *device.ConnectMsg) error
		Disconnected(out *device.ConnectMsg) error
	}
)

func NewInnerLink(conf config.InnerLinkConf) (InnerLink, error) {
	return NewNatsClient(conf.Nats)
}
