// Code generated by goctl. DO NOT EDIT!
// Source: dc.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dcsvr/internal/logic"
	"github.com/i-Things/things/src/dcsvr/internal/svc"
	"github.com/i-Things/things/src/dcsvr/pb/dc"
)

type DcServer struct {
	svcCtx *svc.ServiceContext
	dc.UnimplementedDcServer
}

func NewDcServer(svcCtx *svc.ServiceContext) *DcServer {
	return &DcServer{
		svcCtx: svcCtx,
	}
}

// 管理组
func (s *DcServer) ManageGroupInfo(ctx context.Context, in *dc.ManageGroupInfoReq) (*dc.GroupInfo, error) {
	l := logic.NewManageGroupInfoLogic(ctx, s.svcCtx)
	return l.ManageGroupInfo(in)
}

// 管理组成员
func (s *DcServer) ManageGroupMember(ctx context.Context, in *dc.ManageGroupMemberReq) (*dc.GroupMember, error) {
	l := logic.NewManageGroupMemberLogic(ctx, s.svcCtx)
	return l.ManageGroupMember(in)
}

// 获取组信息
func (s *DcServer) GetGroupInfo(ctx context.Context, in *dc.GetGroupInfoReq) (*dc.GetGroupInfoResp, error) {
	l := logic.NewGetGroupInfoLogic(ctx, s.svcCtx)
	return l.GetGroupInfo(in)
}

// 获取组成员
func (s *DcServer) GetGroupMember(ctx context.Context, in *dc.GetGroupMemberReq) (*dc.GetGroupMemberResp, error) {
	l := logic.NewGetGroupMemberLogic(ctx, s.svcCtx)
	return l.GetGroupMember(in)
}

// 同步调用设备行为
func (s *DcServer) SendAction(ctx context.Context, in *dc.SendActionReq) (*dc.SendActionResp, error) {
	l := logic.NewSendActionLogic(ctx, s.svcCtx)
	return l.SendAction(in)
}

// 同步调用设备属性
func (s *DcServer) SendProperty(ctx context.Context, in *dc.SendPropertyReq) (*dc.SendPropertyResp, error) {
	l := logic.NewSendPropertyLogic(ctx, s.svcCtx)
	return l.SendProperty(in)
}
