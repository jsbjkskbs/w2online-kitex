package main

import (
	"context"
	interact "work/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// LikeAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) LikeAction(ctx context.Context, request *interact.LikeActionRequest) (resp *interact.LikeActionResponse, err error) {
	// TODO: Your code here...
	return
}

// LikeList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) LikeList(ctx context.Context, request *interact.LikeListRequest) (resp *interact.LikeListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentPublish implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentPublish(ctx context.Context, request *interact.CommentPublishRequest) (resp *interact.CommentPublishResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, request *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentDelete implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentDelete(ctx context.Context, request *interact.CommentDeleteRequest) (resp *interact.CommentDeleteResponse, err error) {
	// TODO: Your code here...
	return
}
