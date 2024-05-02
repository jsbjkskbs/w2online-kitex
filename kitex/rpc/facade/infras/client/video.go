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

	suite, closer := jaeger_suite.NewClientTracer().Init(conf.FacadeServiceName)
	defer closer.Close()

	c, err := videoservice.NewClient(
		conf.VideoServiceName,
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(suite),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.FacadeServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func VideoFeed(ctx context.Context, req *video.VideoFeedRequest) (*video.VideoFeedResponseData, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func VideoList(ctx context.Context, req *video.VideoListRequest) (*video.VideoListResponseData, error) {
	resp, err := videoClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func VideoPopular(ctx context.Context, req *video.VideoPopularRequest) (*video.VideoPopularResponseData, error) {
	resp, err := videoClient.Popular(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func VideoSearch(ctx context.Context, req *video.VideoSearchRequest) (*video.VideoSearchResponseData, error) {
	resp, err := videoClient.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func VideoPublishStart(ctx context.Context, req *video.VideoPublishStartRequest) (string, error) {
	resp, err := videoClient.VideoPublishStart(ctx, req)
	if err != nil {
		return ``, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return ``, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Uuid, nil
}

func VideoPublishUploading(ctx context.Context, req *video.VideoPublishUploadingRequest) error {
	resp, err := videoClient.VideoPublishUploading(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func VideoPublishCancle(ctx context.Context, req *video.VideoPublishCancleRequest) error {
	resp, err := videoClient.VideoPublishCancle(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func VideoPublishComplete(ctx context.Context, req *video.VideoPublishCompleteRequest) error {
	resp, err := videoClient.VideoPublishComplete(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func VideoVisit(ctx context.Context, req *video.VideoVisitRequest) (*base.Video, error) {
	resp, err := videoClient.Visit(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Item, nil
}
