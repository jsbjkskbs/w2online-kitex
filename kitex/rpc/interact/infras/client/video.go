package client

import (
	"context"
	"time"
	"work/kitex_gen/base"
	"work/kitex_gen/video"
	"work/kitex_gen/video/videoservice"
	"work/pkg/errno"
	"work/pkg/jaeger_suite"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewClientTracer().Init(conf.InteractServiceName)
	defer closer.Close()

	c, err := videoservice.NewClient(
		conf.VideoServiceName,
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(suite),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.InteractServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func VideoInfo(ctx context.Context, req *video.VideoInfoRequest) (*base.Video, error) {
	resp, err := videoClient.Info(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data.Item, nil
}

func VideoDelete(ctx context.Context, req *video.VideoDeleteRequest) error {
	resp, err := videoClient.Delete(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}