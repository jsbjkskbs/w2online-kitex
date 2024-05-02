package main

import (
	"context"
	"work/kitex_gen/base"
	relation "work/kitex_gen/relation"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/relation/service"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, request *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.RelationActionResponse)

	err = service.NewRelationService(ctx).NewRelationActionEvent(request)
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

// FollowingList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowingList(ctx context.Context, request *relation.FollowingListRequest) (resp *relation.FollowingListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.FollowingListResponse)
	resp.Data = &relation.FollowingListResponseData{}

	data, err := service.NewRelationService(ctx).NewFollowingListEvent(request)
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

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, request *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.FollowerListResponse)
	resp.Data = &relation.FollowerListResponseData{}

	data, err := service.NewRelationService(ctx).NewFollowerEvent(request)
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

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, request *relation.FriendListRequest) (resp *relation.FriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.FriendListResponse)
	resp.Data = &relation.FriendListResponseData{}

	data, err := service.NewRelationService(ctx).NewFriendListEvent(request)
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
