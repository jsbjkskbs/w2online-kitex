package main

import (
	"net"
	video "work/kitex_gen/video/videoservice"
	"work/pkg/jaeger_suite"
	conf "work/rpc/rpc_conf"
	"work/rpc/video/common/conf_loader"
	"work/rpc/video/common/dustman"
	"work/rpc/video/infras/client"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	conf_loader.Init()
	client.Init()
	dustman.NewFileDustman().Run()
	dustman.NewRedisDustman().Run()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.VideoServiceAddress)
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewServerSuite().Init(conf.VideoServiceName)
	defer closer.Close()

	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.VideoServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(suite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.VideoServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}

}
