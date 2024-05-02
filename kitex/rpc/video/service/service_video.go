package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
	"work/kitex_gen/base"
	kuser "work/kitex_gen/user"
	"work/kitex_gen/video"
	"work/pkg/constants"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/video/dal/db"
	"work/rpc/video/infras/client"
	"work/rpc/video/infras/elasticsearch"
	"work/rpc/video/infras/oss"
	"work/rpc/video/infras/redis"
)

type VideoService struct {
	ctx context.Context
}

var (
	TempVideoFolderPath string
)

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{
		ctx: ctx,
	}
}

func (service VideoService) NewCancleUploadEvent(request *video.VideoPublishCancleRequest) error {
	if request.Uuid == `` {
		return errno.RequestError
	}
	if err := service.deleteTempDir(TempVideoFolderPath + request.UserId + `_` + request.Uuid); err != nil {
		return errno.ServiceError
	}
	if err := redis.DeleteVideoEvent(request.Uuid, request.UserId); err != nil {
		return errno.RedisError
	}
	return nil
}

func (service VideoService) NewUploadCompleteEvent(request *video.VideoPublishCompleteRequest) error {
	if request.Uuid == `` {
		return errno.RequestError
	}

	reallyComplete, err := redis.IsChunkAllRecorded(request.Uuid, request.UserId)
	if err != nil {
		return errno.RedisError
	}
	if !reallyComplete {
		return errno.RequestError
	}

	m3u8name, err := redis.GetM3U8Filename(request.Uuid, request.UserId)
	if err != nil {
		return errno.RedisError
	}

	err = utils.M3u8ToMp4(TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+m3u8name,
		TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+`video.mp4`)
	if err != nil {
		return errno.ServiceError
	}

	err = utils.GenerateMp4CoverJpg(TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+`video.mp4`,
		TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+`cover.jpg`)
	if err != nil {
		return errno.ServiceError
	}

	info, err := redis.FinishVideoEvent(request.Uuid, request.UserId)
	if err != nil {
		return errno.RedisError
	}

	uidInt64, _ := strconv.Atoi(request.UserId)
	d := db.Video{
		Title:       info[0],
		Description: info[1],
		UserId:      int64(uidInt64),
		VisitCount:  0,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		DeletedAt:   0,
	}
	vid, err := db.InsertVideo(&d)
	if err != nil {
		return errno.ServiceError
	}

	var (
		videoUrl, coverUrl string
		user               *base.User
		wg                 sync.WaitGroup
	)
	errChan := make(chan error, 3)
	wg.Add(3)
	go func() {
		videoUrl, err = oss.UploadVideo(TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+`video.mp4`, vid)
		if err != nil {
			errChan <- errno.OSSError
		}
		wg.Done()
	}()
	go func() {
		coverUrl, err = oss.UploadVideoCover(TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+`cover.jpg`, vid)
		if err != nil {
			errChan <- errno.OSSError
		}
		wg.Done()
	}()
	go func() {
		user, err = client.UserInfo(service.ctx, &kuser.UserInfoRequest{UserId: request.UserId})
		if err != nil {
			errChan <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		if err := db.DeleteVideo(vid); err != nil {
			return errno.MySQLError
		}
		return err
	default:
	}

	err = db.UpdateVideoUrl(videoUrl, coverUrl, vid)
	if err != nil {
		return errno.MySQLError
	}

	err = elasticsearch.CreateVideoDoc(&elasticsearch.Video{
		Title:       d.Title,
		Description: d.Description,
		Username:    user.Username,
		UserId:      request.UserId,
		CreatedAt:   d.CreatedAt,
		Info: elasticsearch.VideoOtherdata{
			Id:           vid,
			VideoUrl:     videoUrl,
			CoverUrl:     coverUrl,
			UpdatedAt:    d.UpdatedAt,
			DeletedAt:    d.DeletedAt,
			VisitCount:   d.VisitCount,
			LikeCount:    0,
			CommentCount: 0,
		},
	})
	if err != nil {
		return errno.ElasticError
	}

	err = redis.DeleteVideoEvent(request.Uuid, request.UserId)
	if err != nil {
		return errno.RedisError
	}

	err = service.deleteTempDir(TempVideoFolderPath + request.UserId + `_` + request.Uuid)
	if err != nil {
		return errno.ServiceError
	}

	err = redis.PutVideoVisitInfo(vid, `0`)
	if err != nil {
		return errno.RedisError
	}
	err = redis.PutVideoLikeInfo(vid, nil)
	if err != nil {
		return errno.RedisError
	}

	return nil
}

func (service VideoService) NewUploadingEvent(request *video.VideoPublishUploadingRequest) error {
	if request.Filename == `` || request.Uuid == `` || request.ChunkNumber <= 0 {
		return errno.RequestError
	}
	data := request.Data

	if !service.isMD5Same(data, request.Md5) {
		return errno.DataProcessFailed
	}

	if request.IsM3u8 {
		err := redis.RecordM3U8Filename(request.Uuid, request.UserId, request.Filename)
		if err != nil {
			return errno.RedisError
		}
	}

	err := service.saveTempData(TempVideoFolderPath+request.UserId+`_`+request.Uuid+`/`+request.Filename, data)
	if err != nil {
		return errno.ServiceError
	}

	if err := redis.DoneChunkEvent(request.Uuid, request.UserId, request.ChunkNumber); err != nil {
		return errno.RedisError
	}
	return nil
}

func (service VideoService) NewUploadEvent(request *video.VideoPublishStartRequest) (string, error) {
	var (
		uuid = ``
		uid  = request.UserId
		err  error
	)

	if request.Title == `` || request.ChunkTotalNumber <= 0 {
		return ``, errno.RequestError
	}
	uuid, err = redis.NewVideoEvent(request.Title, request.Description, uid, fmt.Sprint(request.ChunkTotalNumber))
	if err != nil {
		return ``, errno.RedisError
	}
	if err := os.Mkdir(TempVideoFolderPath+uid+`_`+uuid, os.ModePerm); err != nil {
		if err := redis.DeleteVideoEvent(uuid, uid); err != nil {
			return ``, errno.RedisError
		}
		return ``, errno.ServiceError
	}
	return uuid, nil
}

func (service VideoService) NewSearchEvent(request *video.VideoSearchRequest) (*video.VideoSearchResponseData, error) {
	items, total, err := elasticsearch.SearchVideoDoc(
		request.Keywords,
		request.Username,
		request.PageSize, request.PageNum,
		request.FromDate, request.ToDate,
	)
	if err != nil {
		return nil, errno.ElasticError
	}
	return &video.VideoSearchResponseData{Items: items, Total: total}, nil
}

func (service VideoService) NewFeedEvent(request *video.VideoFeedRequest) (*video.VideoFeedResponseData, error) {
	var timestamp int64
	if len(request.LatestTime) == 0 {
		timestamp = 0
	} else {
		timestamp, _ = strconv.ParseInt(request.LatestTime, 10, 64)
	}
	items, _, err := elasticsearch.RandomVideoDoc(timestamp)
	if err != nil {
		return nil, errno.ElasticError
	}
	return &video.VideoFeedResponseData{Items: items}, err
}

func (service VideoService) NewListEvent(request *video.VideoListRequest) (*video.VideoListResponseData, error) {
	items, total, err := elasticsearch.SearchVideoDocByUserId(request.UserId, request.PageNum, request.PageSize)
	if err != nil {
		return nil, errno.ElasticError
	}
	return &video.VideoListResponseData{Data: items, Total: total}, nil
}

func (service VideoService) NewVisitEvent(request *video.VideoVisitRequest) (*base.Video, error) {
	vid := request.VideoId
	if err := redis.IncrVideoVisitInfo(vid); err != nil {
		return nil, errno.RedisError
	}
	info, err := elasticsearch.GetVideoDoc(vid)
	if err != nil {
		return nil, errno.ElasticError
	}
	return info, nil
}

func (service VideoService) NewPopularEvent(request *video.VideoPopularRequest) (*video.VideoPopularResponseData, error) {
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = constants.DefaultPageSize
	}
	list, err := redis.GetVideoPopularList(request.PageNum, request.PageSize)
	if err != nil {
		return nil, err
	}
	items := make([]*base.Video, len(*list))
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, len(*list))
	)
	wg.Add(len(*list))
	for i, item := range *list {
		go func(i int, item string) {
			items[i], err = elasticsearch.GetVideoDoc(item)
			if err != nil {
				errChan <- errno.ElasticError
			}
			wg.Done()
		}(i, item)
	}
	wg.Wait()
	select {
	case err := <-errChan:
		return nil, err
	default:
	}
	return &video.VideoPopularResponseData{Items: items}, nil
}

func (service VideoService) NewInfoEvent(request *video.VideoInfoRequest) (*video.VideoInfoResponseData, error) {
	data, err := elasticsearch.GetVideoDoc(request.VideoId)
	if err != nil {
		return nil, errno.ElasticError
	}
	return &video.VideoInfoResponseData{Item: data}, nil
}

func (service VideoService) NewDeleteEvent(request *video.VideoDeleteRequest) error {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 3)
	)
	wg.Add(3)
	go func() {
		if err := db.DeleteVideo(request.VideoId); err != nil {
			errChan <- errno.ServiceError
		}
		wg.Done()
	}()
	go func() {
		if err := redis.DeleteVideo(request.VideoId); err != nil {
			errChan <- errno.RedisError
		}
		wg.Done()
	}()
	go func() {
		if err := elasticsearch.DeleteVideoDoc(request.VideoId); err != nil {
			errChan <- errno.ElasticError
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

func (service VideoService) deleteTempDir(path string) error {
	return os.RemoveAll(path)
}

func (service VideoService) saveTempData(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0777)
}

func (service VideoService) isMD5Same(data []byte, md5 string) bool {
	return utils.GetBytesMD5(data) == md5
}
