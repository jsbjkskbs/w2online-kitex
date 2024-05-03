package main

import (
	"context"
	"work/kitex_gen/base"
	video "work/kitex_gen/video"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/video/service"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, request *video.VideoFeedRequest) (resp *video.VideoFeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoFeedResponse)
	resp.Data = &video.VideoFeedResponseData{}

	data, err := service.NewVideoService(ctx).NewFeedEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data
	return resp, nil
}

// VideoPublishStart implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishStart(ctx context.Context, request *video.VideoPublishStartRequest) (resp *video.VideoPublishStartResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoPublishStartResponse)
	resp.Uuid = ``

	uuid, err := service.NewVideoService(ctx).NewUploadEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Uuid = uuid
	return resp, nil
}

// VideoPublishUploading implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishUploading(ctx context.Context, request *video.VideoPublishUploadingRequest) (resp *video.VideoPublishUploadingResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoPublishUploadingResponse)

	err = service.NewVideoService(ctx).NewUploadingEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// VideoPublishComplete implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishComplete(ctx context.Context, request *video.VideoPublishCompleteRequest) (resp *video.VideoPublishCompleteResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoPublishCompleteResponse)

	err = service.NewVideoService(ctx).NewUploadCompleteEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// VideoPublishCancle implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishCancle(ctx context.Context, request *video.VideoPublishCancleRequest) (resp *video.VideoPublishCancleResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoPublishCancleResponse)

	err = service.NewVideoService(ctx).NewCancleUploadEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// List implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) List(ctx context.Context, request *video.VideoListRequest) (resp *video.VideoListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoListResponse)
	resp.Data = &video.VideoListResponseData{}

	data, err := service.NewVideoService(ctx).NewListEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data
	return resp, nil
}

// Popular implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Popular(ctx context.Context, request *video.VideoPopularRequest) (resp *video.VideoPopularResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoPopularResponse)
	resp.Data = &video.VideoPopularResponseData{}

	data, err := service.NewVideoService(ctx).NewPopularEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data
	return resp, nil
}

// Search implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Search(ctx context.Context, request *video.VideoSearchRequest) (resp *video.VideoSearchResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoSearchResponse)
	resp.Data = &video.VideoSearchResponseData{}

	data, err := service.NewVideoService(ctx).NewSearchEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data
	return resp, nil
}

// Info implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Info(ctx context.Context, request *video.VideoInfoRequest) (resp *video.VideoInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoInfoResponse)
	resp.Data = &video.VideoInfoResponseData{}

	data, err := service.NewVideoService(ctx).NewInfoEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		resp.Data.Item = &base.Video{}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data
	return resp, nil
}

// Delete implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Delete(ctx context.Context, request *video.VideoDeleteRequest) (resp *video.VideoDeleteResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoDeleteResponse)

	err = service.NewVideoService(ctx).NewDeleteEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// IdList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IdList(ctx context.Context, request *video.VideoIdListRequest) (resp *video.VideoIdListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoIdListResponse)

	isEnd, list, err := service.NewVideoService(ctx).NewIdListEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.IsEnd = isEnd
	resp.List = *list
	return resp, nil
}

// UpdateVisitCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVisitCount(ctx context.Context, request *video.UpdateVisitCountRequest) (resp *video.UpdateVisitCountResponse, err error) {
	// TODO: Your code here...
	resp = new(video.UpdateVisitCountResponse)

	err = service.NewVideoService(ctx).NewUpdateVideoVisitCountEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// GetVideoVisitCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoVisitCount(ctx context.Context, request *video.GetVideoVisitCountRequest) (resp *video.GetVideoVisitCountResponse, err error) {
	// TODO: Your code here...
	resp = new(video.GetVideoVisitCountResponse)

	count, err := service.NewVideoService(ctx).NewGetVisitCountEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.VisitCount = count
	return resp, nil
}
