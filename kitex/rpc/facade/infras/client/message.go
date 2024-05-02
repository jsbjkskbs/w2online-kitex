package client

import (
	"context"
	"time"
	"work/kitex_gen/message"
	"work/kitex_gen/message/messageservice"
	"work/pkg/errno"
	"work/pkg/jaeger_suite"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessageRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewClientTracer().Init(conf.FacadeServiceName)
	defer closer.Close()

	c, err := messageservice.NewClient(
		conf.MessageServiceName,
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
	messageClient = c
}

func InsertMessage(ctx context.Context, req *message.InsertMessageRequest) error {
	resp, err := messageClient.InsertMessage(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func PopMessage(ctx context.Context, req *message.PopMessageRequest) (*message.PopMessageResponseData, error) {
	resp, err := messageClient.PopMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}
