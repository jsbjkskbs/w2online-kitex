package main

import (
	"context"
	"work/kitex_gen/base"
	interact "work/kitex_gen/interact"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/interact/service"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// LikeAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) LikeAction(ctx context.Context, request *interact.LikeActionRequest) (resp *interact.LikeActionResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.LikeActionResponse)

	err = service.NewInteractService(ctx).NewLikeActionEvent(request)
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

// LikeList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) LikeList(ctx context.Context, request *interact.LikeListRequest) (resp *interact.LikeListResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.LikeListResponse)
	resp.Data = &interact.LikeListResponseData{}

	data, err := service.NewInteractService(ctx).NewLikeListEvent(request)
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

// CommentPublish implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentPublish(ctx context.Context, request *interact.CommentPublishRequest) (resp *interact.CommentPublishResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.CommentPublishResponse)

	err = service.NewInteractService(ctx).NewCommentPublishEvent(request)
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

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, request *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.CommentListResponse)
	resp.Data = &interact.CommentListResponseData{}

	data, err := service.NewInteractService(ctx).NewCommentListEvent(request)
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

// CommentDelete implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentDelete(ctx context.Context, request *interact.CommentDeleteRequest) (resp *interact.CommentDeleteResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.CommentDeleteResponse)

	err = service.NewInteractService(ctx).NewDeleteEvent(request)
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

// VideoPopularList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) VideoPopularList(ctx context.Context, request *interact.VideoPopularListRequest) (resp *interact.VideoPopularListResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.VideoPopularListResponse)
	resp.Data = &interact.VideoPopularListResponseData{}
	resp.Data.List = make([]string, 0)

	data, err := service.NewInteractService(ctx).NewVideoPopularListEvent(request)
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
	resp.Data.List = *data
	return resp, nil
}

// DeleteVideoInfo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteVideoInfo(ctx context.Context, request *interact.DeleteVideoInfoRequest) (resp *interact.DeleteVideoInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.DeleteVideoInfoResponse)

	err = service.NewInteractService(ctx).NewDeleteVideoInfoEvent(request)
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
