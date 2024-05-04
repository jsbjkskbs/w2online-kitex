package main

import (
	"net"
	user "work/kitex_gen/user/userservice"
	conf "work/rpc/rpc_conf"
	"work/rpc/user/common/conf_loader"

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
	//pprof.Load()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.UserServiceAddress)
	if err != nil {
		panic(err)
	}

	//suite, closer := jaeger_suite.NewServerSuite().Init(conf.UserServiceName)
	//defer closer.Close()

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		//server.WithSuite(suite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserServiceName}),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
