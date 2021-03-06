package main

import (
	"flag"
	"fmt"

	"github.com/i-Things/things/src/dcsvr/internal/config"
	"github.com/i-Things/things/src/dcsvr/internal/server"
	"github.com/i-Things/things/src/dcsvr/internal/svc"
	"github.com/i-Things/things/src/dcsvr/pb/dc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/dc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewDcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		dc.RegisterDcServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
