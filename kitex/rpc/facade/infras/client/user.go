package client

import (
	"context"
	"time"
	"work/kitex_gen/base"
	"work/kitex_gen/user"
	"work/kitex_gen/user/userservice"
	"work/pkg/errno"
	"work/pkg/jaeger_suite"
	conf "work/rpc/rpc_conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger_suite.NewClientTracer().Init(conf.FacadeServiceName)
	defer closer.Close()

	c, err := userservice.NewClient(
		conf.UserServiceName,
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
	userClient = c
}

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (*base.User, error) {
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func UserRegister(ctx context.Context, req *user.UserRegisterRequest) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*base.User, error) {
	resp, err := userClient.Info(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func UserAvatarUpload(ctx context.Context, req *user.UserAvatarUploadRequest) (*base.User, error) {
	resp, err := userClient.AvatarUpload(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func AuthMfaQrcode(ctx context.Context, req *user.AuthMfaQrcodeRequest) (*user.Qrcode, error) {
	resp, err := userClient.AuthMfaQrcode(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != errno.NoError.Code {
		return nil, errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Data, nil
}

func AuthMfaBind(ctx context.Context, req *user.AuthMfaBindRequest) error {
	resp, err := userClient.AuthMfaBind(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != errno.NoError.Code {
		return errno.NewErrorMessage(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}
