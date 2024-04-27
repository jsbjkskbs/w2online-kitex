package client

import (
	"context"
	"time"
	"work/kitex_gen/interact"
	"work/kitex_gen/interact/interactservice"
	"work/pkg/errmsg"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var interactClient interactservice.Client

func initInteractRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := interactservice.NewClient(
		conf.InteractServiceName,
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
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
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func CommentList(ctx context.Context, req *interact.CommentListRequest) (*interact.CommentListResponseData, error) {
	resp, err := interactClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return nil, errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func CommentPublish(ctx context.Context, req *interact.CommentPublishRequest) error {
	resp, err := interactClient.CommentPublish(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func LikeAction(ctx context.Context, req *interact.LikeActionRequest) error {
	resp, err := interactClient.LikeAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func LikeList(ctx context.Context, req *interact.LikeListRequest) (*interact.LikeListResponseData, error) {
	resp, err := interactClient.LikeList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return nil, errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}
