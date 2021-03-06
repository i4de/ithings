package tests

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/src/dcsvr/dc"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

func TestManageGroupInfo(t *testing.T) {
	fmt.Println("ManageGroupInfo")
	client := dc.NewDc(zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: discov.EtcdConf{
		Hosts: []string{"127.0.0.1:2379"},
		Key:   "dc.rpc",
	}}))
	ctx := context.Background()
	{
		req := &dc.ManageGroupInfoReq{
			Opt: def.OPT_ADD,
			Info: &dc.GroupInfo{
				Name: "测试组1",              //组名
				Uid:  1699809227385606144, //管理员用户id
			},
		}
		info, err := client.ManageGroupInfo(ctx, req)
		t.Log(req, info)
		if err != nil {
			t.Errorf("%+v", errors.Fmt(err))
		}
		req = &dc.ManageGroupInfoReq{
			Opt: def.OPT_MODIFY,
			Info: &dc.GroupInfo{
				GroupID: info.GroupID,
				Name:    "测试组1修改",            //组名
				Uid:     1699809227385606144, //管理员用户id
			},
		}
		info, err = client.ManageGroupInfo(ctx, req)
		t.Log(req, info)
		if err != nil {
			t.Errorf("%+v", errors.Fmt(err))
		}
	}
	{
		req := &dc.ManageGroupInfoReq{
			Opt: def.OPT_ADD,
			Info: &dc.GroupInfo{
				Name: "测试组2",              //组名
				Uid:  1699809227385606145, //管理员用户id
			},
		}
		info, err := client.ManageGroupInfo(ctx, req)
		t.Log(req, info)
		if err != nil {
			t.Errorf("%+v", errors.Fmt(err))
		}
	}
}
