// Code generated by Kitex v0.9.1. DO NOT EDIT.

package interactservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	interact "work/kitex_gen/interact"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	LikeAction(ctx context.Context, request *interact.LikeActionRequest, callOptions ...callopt.Option) (r *interact.LikeActionResponse, err error)
	LikeList(ctx context.Context, request *interact.LikeListRequest, callOptions ...callopt.Option) (r *interact.LikeListResponse, err error)
	CommentPublish(ctx context.Context, request *interact.CommentPublishRequest, callOptions ...callopt.Option) (r *interact.CommentPublishResponse, err error)
	CommentList(ctx context.Context, request *interact.CommentListRequest, callOptions ...callopt.Option) (r *interact.CommentListResponse, err error)
	CommentDelete(ctx context.Context, request *interact.CommentDeleteRequest, callOptions ...callopt.Option) (r *interact.CommentDeleteResponse, err error)
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
	return &kInteractServiceClient{
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

type kInteractServiceClient struct {
	*kClient
}

func (p *kInteractServiceClient) LikeAction(ctx context.Context, request *interact.LikeActionRequest, callOptions ...callopt.Option) (r *interact.LikeActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeAction(ctx, request)
}

func (p *kInteractServiceClient) LikeList(ctx context.Context, request *interact.LikeListRequest, callOptions ...callopt.Option) (r *interact.LikeListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeList(ctx, request)
}

func (p *kInteractServiceClient) CommentPublish(ctx context.Context, request *interact.CommentPublishRequest, callOptions ...callopt.Option) (r *interact.CommentPublishResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentPublish(ctx, request)
}

func (p *kInteractServiceClient) CommentList(ctx context.Context, request *interact.CommentListRequest, callOptions ...callopt.Option) (r *interact.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, request)
}

func (p *kInteractServiceClient) CommentDelete(ctx context.Context, request *interact.CommentDeleteRequest, callOptions ...callopt.Option) (r *interact.CommentDeleteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentDelete(ctx, request)
}