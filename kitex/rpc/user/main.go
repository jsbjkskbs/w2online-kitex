package main

import (
	"net"
	user "work/kitex_gen/user/userservice"
	conf "work/rpc/rpc_conf"
	"work/rpc/user/conf_loader"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	if err := conf_loader.Run(); err != nil {
		panic(err)
	}
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.UserServiceAddress)
	if err != nil {
		panic(err)
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
