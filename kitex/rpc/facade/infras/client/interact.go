package client

import (
	"context"
	"time"
	"work/kitex_gen/interact"
	"work/kitex_gen/interact/interactservice"
	"work/pkg/errno"
	"work/pkg/jaeger_suite"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var interactClient interactservice.Client

func initInteractRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewClientTracer().Init(conf.FacadeServiceName)
	defer closer.Close()

	c, err := interactservice.NewClient(
		conf.InteractServiceName,
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
	interactClient = c
}

func CommentDelete(ctx context.Context, req *interact.CommentDeleteRequest) error {
	resp, err := interactClient.CommentDelete(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func CommentList(ctx context.Context, req *interact.CommentListRequest) (*interact.CommentListResponseData, error) {
	resp, err := interactClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func CommentPublish(ctx context.Context, req *interact.CommentPublishRequest) error {
	resp, err := interactClient.CommentPublish(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func LikeAction(ctx context.Context, req *interact.LikeActionRequest) error {
	resp, err := interactClient.LikeAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func LikeList(ctx context.Context, req *interact.LikeListRequest) (*interact.LikeListResponseData, error) {
	resp, err := interactClient.LikeList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}
