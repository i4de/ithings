package tests

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/dm"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

func TestLoginAuth(t *testing.T) {
	fmt.Println("TestLoginAuth")
	client := dm.NewDm(zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: discov.EtcdConf{
		Hosts: []string{"127.0.0.1:2379"},
		Key:   "dm.rpc",
	}}))
	ctx := context.Background()
	productID := "21CYs1k9YpG"
	deviceName := "test8"
	clientID := fmt.Sprintf("%s%s", productID, deviceName)
	//生成 MQTT 的 username 部分, 格式为 ${clientid};${sdkappid};${connid};${expiry}
	userName := fmt.Sprintf("%s;12010126;fawef;1822730956", clientID)
	password, _ := base64.StdEncoding.DecodeString("0f0G7gwaxXgj3/hWNz14FAZJw98=")

	pwd := fmt.Sprintf("%s;hmacsha1", utils.HmacSha1(userName, password))
	req := &dm.LoginAuthReq{
		Username:    userName,      //用户名
		Password:    pwd,           //密码
		ClientID:    clientID,      //clientID
		Ip:          "192.168.1.2", //访问的ip地址
		Certificate: []byte{},      //客户端证书
	}
	info, err := client.LoginAuth(ctx, req)
	t.Log(req, info)
	if err != nil {
		t.Errorf("%+v", errors.Fmt(err))
	}
}
