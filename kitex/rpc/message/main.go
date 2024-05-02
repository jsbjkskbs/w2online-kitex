package main

import (
	"net"
	message "work/kitex_gen/message/messageservice"
	"work/pkg/jaeger_suite"
	"work/rpc/message/common/conf_loader"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	conf_loader.Init()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.MessageServiceAddress)
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewServerSuite().Init(conf.MessageServiceName)
	defer closer.Close()

	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.MessageServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(suite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.MessageServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
