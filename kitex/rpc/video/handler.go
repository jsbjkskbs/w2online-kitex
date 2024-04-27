package main

import (
	"context"
	video "work/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, request *video.VideoFeedRequest) (resp *video.VideoFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishStart implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishStart(ctx context.Context, request *video.VideoPublishStartRequest) (resp *video.VideoPublishStartResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishUploading implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishUploading(ctx context.Context, request *video.VideoPublishUploadingRequest) (resp *video.VideoPublishUploadingResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishComplete implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishComplete(ctx context.Context, request *video.VideoPublishCompleteRequest) (resp *video.VideoPublishCompleteResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishCancle implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishCancle(ctx context.Context, request *video.VideoPublishCancleRequest) (resp *video.VideoPublishCancleResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) List(ctx context.Context, request *video.VideoListRequest) (resp *video.VideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// Popular implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Popular(ctx context.Context, request *video.VideoPopularRequest) (resp *video.VideoPopularResponse, err error) {
	// TODO: Your code here...
	return
}

// Search implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Search(ctx context.Context, request *video.VideoSearchRequest) (resp *video.VideoSearchResponse, err error) {
	// TODO: Your code here...
	return
}

// Visit implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Visit(ctx context.Context, request *video.VideoVisitRequest) (resp *video.VideoVisitResponse, err error) {
	// TODO: Your code here...
	return
}
