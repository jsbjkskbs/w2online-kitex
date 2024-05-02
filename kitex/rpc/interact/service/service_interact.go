package service

import (
	"context"
	"fmt"
	"sync"
	"time"
	"work/kitex_gen/base"
	"work/kitex_gen/interact"
	"work/kitex_gen/video"
	"work/pkg/constants"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/interact/dal/db"
	"work/rpc/interact/infras/client"
	"work/rpc/interact/infras/rabbitmq"
	"work/rpc/interact/infras/redis"
)

type InteractService struct {
	ctx context.Context
}

func NewInteractService(ctx context.Context) *InteractService {
	return &InteractService{ctx: ctx}
}

func (service InteractService) NewCommentPublishEvent(request *interact.CommentPublishRequest) error {
	uid := request.UserId
	if request.Content == `` {
		return errno.RequestError
	}
	if request.CommentId == `` && request.VideoId == `` {
		return errno.RequestError
	}
	if request.CommentId == `` {
		request.CommentId = `-1`
	} else {
		parentCommentId, err := db.GetParentCommentId(request.CommentId)
		if err != nil {
			return errno.ServiceError
		}
		if parentCommentId != `-1` {
			request.CommentId = parentCommentId
		}
	}
	if request.VideoId == `` {
		vid, err := db.GetCommentVideoId(request.CommentId)
		if err != nil {
			return errno.ServiceError
		}
		request.VideoId = vid
	}

	newComment := db.Comment{
		VideoId:   request.VideoId,
		ParentId:  request.CommentId,
		UserId:    uid,
		Content:   request.Content,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		DeletedAt: 0,
	}

	if err := rabbitmq.CommentMQ.Send(&newComment); err != nil {
		return err
	}
	return nil
}

func (service InteractService) NewLikeActionEvent(request *interact.LikeActionRequest) error {
	uid := request.UserId
	if request.VideoId != `` {
		switch request.ActionType {
		case `1`:
			{
				if err := redis.AppendVideoLikeInfo(request.VideoId, uid); err != nil {
					return errno.RedisError
				}
			}
		case `2`:
			{
				if err := redis.RemoveVideoLikeInfo(request.VideoId, uid); err != nil {
					return errno.RedisError
				}
			}
		}
	} else if request.CommentId != `` {
		switch request.ActionType {
		case `1`:
			{
				if err := redis.AppendCommentLikeInfo(request.CommentId, uid); err != nil {
					return errno.RedisError
				}
			}
		case `2`:
			{
				if err := redis.RemoveCommentLikeInfo(request.CommentId, uid); err != nil {
					return errno.RedisError
				}
			}
		}
	} else {
		return errno.RequestError
	}
	return nil
}

func (service InteractService) NewLikeListEvent(request *interact.LikeListRequest) (*interact.LikeListResponseData, error) {
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = constants.DefaultPageSize
	}
	list, err := db.GetVideoLikeListByUserId(request.UserId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ServiceError
	}
	data := make([]*base.Video, len(*list))
	for i, item := range *list {
		if data[i], err = client.VideoInfo(context.Background(), &video.VideoInfoRequest{VideoId: item}); err != nil {
			return nil, errno.ElasticError
		}
	}
	return &interact.LikeListResponseData{Items: data}, nil
}

func (service InteractService) NewCommentListEvent(request *interact.CommentListRequest) (*interact.CommentListResponseData, error) {
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = constants.DefaultPageSize
	}

	var (
		data *[]*base.Comment
		err  error
	)
	if request.VideoId != `` {
		if data, err = getVideoComment(request); err != nil {
			return nil, err
		}
	} else if request.CommentId != `` {
		if data, err = getCommentComment(request); err != nil {
			return nil, err
		}
	} else {
		return nil, errno.RequestError
	}
	return &interact.CommentListResponseData{Items: *data}, nil
}

func (service InteractService) NewDeleteEvent(request *interact.CommentDeleteRequest) error {
	if request.VideoId != `` {
		videoInfo, err := client.VideoInfo(context.Background(), &video.VideoInfoRequest{VideoId: request.VideoId})
		if err != nil {
			return errno.RPCError
		}
		if videoInfo.UserId != request.FromUserId {
			return errno.ServiceError
		}
		if err := deleteVideo(request); err != nil {
			return err
		}
	} else if request.CommentId != `` {
		commentInfo, err := db.GetCommentInfo(request.CommentId)
		if err != nil {
			return errno.MySQLError
		}
		if commentInfo.UserId != request.FromUserId {
			return errno.ServiceError
		}
		if err := deleteComment(request); err != nil {
			return err
		}
	} else {
		return errno.RequestError
	}
	return nil
}

func getVideoComment(request *interact.CommentListRequest) (*[]*base.Comment, error) {
	data := make([]*base.Comment, 0)
	list, err := db.GetVideoCommentListByPart(request.VideoId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ServiceError
	}
	for _, item := range *list {
		d, err := db.GetCommentInfo(item)
		if err != nil {
			return nil, errno.ServiceError
		}
		likeCount, err := db.GetCommentLikeCount(item)
		if err != nil {
			return nil, errno.ServiceError
		}
		childCount, err := db.GetChildCommentCount(item)
		if err != nil {
			return nil, errno.ServiceError
		}
		data = append(data, &base.Comment{
			Id:         fmt.Sprint(d.Id),
			UserId:     d.UserId,
			VideoId:    d.VideoId,
			ParentId:   d.ParentId,
			LikeCount:  likeCount,
			ChildCount: childCount,
			Content:    d.Content,
			CreatedAt:  utils.ConvertTimestampToStringDefault(d.CreatedAt),
			UpdatedAt:  utils.ConvertTimestampToStringDefault(d.UpdatedAt),
			DeletedAt:  utils.ConvertTimestampToStringDefault(d.DeletedAt),
		})
	}
	return &data, nil
}

func getCommentComment(request *interact.CommentListRequest) (*[]*base.Comment, error) {
	data := make([]*base.Comment, 0)
	list, err := db.GetCommentChildListByPart(request.CommentId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ServiceError
	}
	for _, item := range *list {
		d, err := db.GetCommentInfo(item)
		if err != nil {
			return nil, errno.ServiceError
		}
		likeCount, err := db.GetCommentLikeCount(item)
		if err != nil {
			return nil, errno.ServiceError
		}
		childCount, err := db.GetChildCommentCount(item)
		if err != nil {
			return nil, errno.ServiceError
		}
		data = append(data, &base.Comment{
			Id:         fmt.Sprint(d.Id),
			UserId:     d.UserId,
			VideoId:    d.VideoId,
			ParentId:   d.ParentId,
			LikeCount:  likeCount,
			ChildCount: childCount,
			CreatedAt:  utils.ConvertTimestampToStringDefault(d.CreatedAt),
			UpdatedAt:  utils.ConvertTimestampToStringDefault(d.UpdatedAt),
			DeletedAt:  utils.ConvertTimestampToStringDefault(d.DeletedAt),
		})
	}
	return &data, nil
}

func deleteVideo(request *interact.CommentDeleteRequest) error {
	list, err := db.GetVideoCommentList(request.VideoId)
	if err != nil {
		return errno.MySQLError
	}
	if err := client.VideoDelete(context.Background(), &video.VideoDeleteRequest{VideoId: request.VideoId}); err != nil {
		return errno.ServiceError
	}

	var (
		wg      sync.WaitGroup
		errChan = make(chan error, len(*list))
	)
	wg.Add(len(*list))
	for _, item := range *list {
		go func(cid string) {
			if err := deleteComment(&interact.CommentDeleteRequest{CommentId: cid}); err != nil {
				errChan <- err
			}
			wg.Done()
		}(item)
	}
	wg.Wait()
	select {
	case result := <-errChan:
		return result
	default:
	}
	return nil
}

func deleteComment(request *interact.CommentDeleteRequest) error {
	if err := db.DeleteComment(request.CommentId); err != nil {
		return errno.ServiceError
	}
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 2)
	)
	wg.Add(2)
	go func() {
		if err := db.DeleteComment(request.CommentId); err != nil {
			errChan <- errno.RedisError
		}
		wg.Done()
	}()
	go func() {
		if err := redis.DeleteCommentAndAllAbout(request.CommentId); err != nil {
			errChan <- errno.RedisError
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		return err
	default:
	}
	return nil
}
