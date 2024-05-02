package client

import (
	"context"
	"time"
	"work/kitex_gen/relation"
	"work/kitex_gen/relation/relationservice"
	"work/pkg/errno"
	"work/pkg/jaeger_suite"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewClientTracer().Init(conf.FacadeServiceName)
	defer closer.Close()

	c, err := relationservice.NewClient(
		conf.RelationServiceName,
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
	relationClient = c
}

func FollowerList(ctx context.Context, req *relation.FollowerListRequest) (*relation.FollowerListResponseData, error) {
	resp, err := relationClient.FollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func FollowingList(ctx context.Context, req *relation.FollowingListRequest) (*relation.FollowingListResponseData, error) {
	resp, err := relationClient.FollowingList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func FriendList(ctx context.Context, req *relation.FriendListRequest) (*relation.FriendListResponseData, error) {
	resp, err := relationClient.FriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) error {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}
