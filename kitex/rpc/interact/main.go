package main

import (
	"net"
	interact "work/kitex_gen/interact/interactservice"
	"work/pkg/jaeger_suite"
	"work/rpc/interact/common/conf_loader"
	"work/rpc/interact/common/syncman"
	"work/rpc/interact/infras/client"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	conf_loader.Init()
	syncman.NewCommentSyncman().Run()
	syncman.NewVideoSyncman().Run()
	client.Init()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.InteractServiceAddress)
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewServerSuite().Init(conf.InteractServiceName)
	defer closer.Close()

	svr := interact.NewServer(new(InteractServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.InteractServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(suite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.InteractServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
