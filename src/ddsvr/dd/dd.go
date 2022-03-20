// Code generated by goctl. DO NOT EDIT!
// Source: dd.proto

package dd

import (
	"context"

	"github.com/i-Things/things/src/ddsvr/pb/dd"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty = dd.Empty

	Dd interface {
		// ping pong 测试
		Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultDd struct {
		cli zrpc.Client
	}
)

func NewDd(cli zrpc.Client) Dd {
	return &defaultDd{
		cli: cli,
	}
}

// ping pong 测试
func (m *defaultDd) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	client := dd.NewDdClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}