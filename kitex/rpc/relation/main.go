package main

import (
	"net"
	relation "work/kitex_gen/relation/relationservice"
	"work/pkg/jaeger_suite"
	"work/rpc/relation/common/conf_loader"
	"work/rpc/relation/infras/client"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	conf_loader.Init()
	client.Init()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.RelationServiceAddress)
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewServerSuite().Init(conf.RelationServiceName)
	defer closer.Close()

	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.RelationServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(suite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.RelationServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}

}
