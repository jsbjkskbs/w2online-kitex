package client

import (
	"context"
	"time"
	"work/kitex_gen/message"
	"work/kitex_gen/message/messageservice"
	"work/pkg/errmsg"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessageRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		conf.MessageServiceName,
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

func InsertMessage(ctx context.Context, req *message.InsertMessageRequest) error {
	resp, err := messageClient.InsertMessage(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func PopMessage(ctx context.Context, req *message.PopMessageRequest) (*message.PopMessageResponseData, error) {
	resp, err := messageClient.PopMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errmsg.NoError.ErrorCode {
		return nil, errmsg.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}
