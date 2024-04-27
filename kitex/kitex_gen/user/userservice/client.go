// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "work/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, request *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error)
	Login(ctx context.Context, request *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error)
	Info(ctx context.Context, request *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error)
	AvatarUpload(ctx context.Context, request *user.UserAvatarUploadRequest, callOptions ...callopt.Option) (r *user.UserAvatarUploadResponse, err error)
	AuthMfaQrcode(ctx context.Context, request *user.AuthMfaQrcodeRequest, callOptions ...callopt.Option) (r *user.AuthMfaQrcodeResponse, err error)
	AuthMfaBind(ctx context.Context, request *user.AuthMfaBindRequest, callOptions ...callopt.Option) (r *user.AuthMfaBindResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, request *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, request)
}

func (p *kUserServiceClient) Login(ctx context.Context, request *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, request)
}

func (p *kUserServiceClient) Info(ctx context.Context, request *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Info(ctx, request)
}

func (p *kUserServiceClient) AvatarUpload(ctx context.Context, request *user.UserAvatarUploadRequest, callOptions ...callopt.Option) (r *user.UserAvatarUploadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AvatarUpload(ctx, request)
}

func (p *kUserServiceClient) AuthMfaQrcode(ctx context.Context, request *user.AuthMfaQrcodeRequest, callOptions ...callopt.Option) (r *user.AuthMfaQrcodeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthMfaQrcode(ctx, request)
}

func (p *kUserServiceClient) AuthMfaBind(ctx context.Context, request *user.AuthMfaBindRequest, callOptions ...callopt.Option) (r *user.AuthMfaBindResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthMfaBind(ctx, request)
}