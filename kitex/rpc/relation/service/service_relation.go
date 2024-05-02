package service

import (
	"context"

	"work/kitex_gen/base"
	"work/kitex_gen/relation"
	"work/kitex_gen/user"
	"work/pkg/constants"
	"work/pkg/errno"
	"work/rpc/relation/dal/db"
	"work/rpc/relation/infras/client"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (service RelationService) NewRelationActionEvent(request *relation.RelationActionRequest) error {
	if request.FromUserId == request.ToUserId {
		return errno.RequestError
	}
	switch request.ActionType {
	case 0:
		if err := createFollow(request.FromUserId, request); err != nil {
			return err
		}
	case 1:
		if err := cancleFollow(request.FromUserId, request); err != nil {
			return err
		}
	}
	return nil
}

func (service RelationService) NewFollowingListEvent(request *relation.FollowingListRequest) (*relation.FollowingListResponseData, error) {
	userInfo, err := client.UserInfo(context.Background(), &user.UserInfoRequest{UserId: request.UserId})
	if err != nil {
		return nil, errno.ServiceError
	}
	if userInfo == nil {
		return nil, errno.InfomationNotExist
	}
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = constants.DefaultPageSize
	}
	list, err := db.GetFollowListPaged(request.UserId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ServiceError
	}
	data := make([]*base.UserLite, 0)
	for _, item := range *list {
		userInfo, err := client.UserInfo(context.Background(), &user.UserInfoRequest{UserId: item})
		if err != nil {
			return nil, errno.ServiceError
		}
		d := base.UserLite{
			Uid:       item,
			Username:  userInfo.Username,
			AvatarUrl: userInfo.AvatarUrl,
		}
		data = append(data, &d)
	}
	total, err := db.GetFollowCount(request.UserId)
	if err != nil {
		return nil, errno.ServiceError
	}
	return &relation.FollowingListResponseData{Items: data, Total: total}, nil
}

func (service RelationService) NewFollowerEvent(request *relation.FollowerListRequest) (*relation.FollowerListResponseData, error) {
	userInfo, err := client.UserInfo(context.Background(), &user.UserInfoRequest{UserId: request.UserId})
	if err != nil {
		return nil, errno.ServiceError
	}
	if userInfo == nil {
		return nil, errno.InfomationNotExist
	}
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = constants.DefaultPageSize
	}
	list, err := db.GetFollowerListPaged(request.UserId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ServiceError
	}
	data := make([]*base.UserLite, 0)
	for _, item := range *list {
		userInfo, err := client.UserInfo(context.Background(), &user.UserInfoRequest{UserId: request.UserId})
		if err != nil {
			return nil, errno.ServiceError
		}
		d := base.UserLite{
			Uid:       item,
			Username:  userInfo.Username,
			AvatarUrl: userInfo.AvatarUrl,
		}
		data = append(data, &d)
	}
	total, err := db.GetFollowerCount(request.UserId)
	if err != nil {
		return nil, errno.ServiceError
	}
	return &relation.FollowerListResponseData{Items: data, Total: total}, nil
}

func (service RelationService) NewFriendListEvent(request *relation.FriendListRequest) (*relation.FriendListResponseData, error) {
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = constants.DefaultPageSize
	}
	list, err := db.GetFriendListPaged(request.UserId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ServiceError
	}
	data := make([]*base.UserLite, 0)
	for _, item := range *list {
		userInfo, err := client.UserInfo(service.ctx, &user.UserInfoRequest{UserId: request.UserId})
		if err != nil {
			return nil, errno.ServiceError
		}
		d := base.UserLite{
			Uid:       item,
			Username:  userInfo.Username,
			AvatarUrl: userInfo.AvatarUrl,
		}
		data = append(data, &d)
	}
	total, err := db.GetFriendCount(request.UserId)
	if err != nil {
		return nil, errno.ServiceError
	}
	return &relation.FriendListResponseData{Items: data, Total: total}, nil

}

func createFollow(uid string, request *relation.RelationActionRequest) error {
	if err := db.CreateFollow(request.ToUserId, uid); err != nil {
		return errno.ServiceError
	}
	return nil
}

func cancleFollow(uid string, request *relation.RelationActionRequest) error {
	if err := db.DeleteFollow(request.ToUserId, uid); err != nil {
		return errno.ServiceError
	}
	return nil
}
