package syncman

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
	"work/rpc/interact/dal/db"
	"work/rpc/interact/infras/elasticsearch"
	"work/rpc/interact/infras/redis"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoSyncman struct {
	ctx    context.Context
	cancle context.CancelFunc
}

func NewVideoSyncman() *VideoSyncman {
	ctx, cancle := context.WithCancel(context.Background())
	return &VideoSyncman{
		ctx:    ctx,
		cancle: cancle,
	}
}

func (sm VideoSyncman) Run() {
	if err := videoSyncMwWhenInit(); err != nil {
		panic(err)
	}
	go func() {
		for {
			time.Sleep(time.Minute * 10)
			select {
			case <-sm.ctx.Done():
				hlog.Info("Ok,stop sync[video]")
				return
			default:
			}
			var (
				wg           sync.WaitGroup
				errChan      = make(chan error, 3)
				commentCount int64
				likeList     *[]string
				vidList      *[]string
				dislikeList  *[]string
			)
			var err error
			if vidList, err = db.GetVideoIdList(); err != nil {
				hlog.Warn(err)
			}
			for _, vid := range *vidList {
				wg.Add(3)
				go func() {
					var err error
					commentCountString, err := db.GetVideoCommentCount(vid)
					if err != nil {
						errChan <- err
					}
					commentCount, _ = strconv.ParseInt(commentCountString, 10, 64)
					wg.Done()
				}()
				go func() {
					var err error
					if likeList, err = redis.GetNewUpdateVideoLikeList(vid); err != nil {
						errChan <- err
					}
					wg.Done()
				}()
				go func() {
					var err error
					if dislikeList, err = redis.GetNewDeleteVideoLikeList(vid); err != nil {
						errChan <- err
					}
					wg.Done()
				}()
				wg.Wait()
				select {
				case result := <-errChan:
					hlog.Error(result)
					continue
				default:
				}
				likeCount, err := redis.GetVideoLikeCount(vid)
				if err != nil {
					hlog.Error(err)
					continue
				}
				for _, uid := range *likeList {
					if err := db.CreateVideoLike(&db.VideoLike{
						UserId:    uid,
						VideoId:   vid,
						CreatedAt: time.Now().Unix(),
						DeletedAt: 0,
					}); err != nil {
						hlog.Error(err)
					}
					if err := redis.AppendVideoLikeInfoToStaticSpace(vid, uid); err != nil {
						hlog.Error(err)
					}
					if err := redis.DeleteVideoLikeInfoFromDynamicSpace(vid, uid); err != nil {
						hlog.Error(err)
					}
				}
				for _, uid := range *dislikeList {
					if err := db.DeleteVideoLike(vid, uid); err != nil {
						hlog.Error(err)
					}
					if err := redis.DeleteVideoLikeInfoFromDynamicSpace(vid, uid); err != nil {
						hlog.Error(err)
					}
				}

				err = elasticsearch.UpdateVideoCommentAndLikeCount(vid, fmt.Sprint(likeCount), fmt.Sprint(commentCount))
				if err != nil {
					hlog.Error(err)
				}
			}
		}
	}()
}

func (sm VideoSyncman) Stop() {
	sm.cancle()
}

type videoSyncData struct {
	vid          string
	likeList     *[]string
	commentCount string
}

func videoSyncMwWhenInit() error {
	list, err := db.GetVideoIdList()
	if err != nil {
		panic(err)
	}

	var (
		wg       sync.WaitGroup
		errChan  = make(chan error, 2)
		syncList = make([]videoSyncData, 0)
		data     videoSyncData
	)
	for _, vid := range *list {
		data.vid = vid
		wg.Add(2)
		go func(data *videoSyncData) {
			if data.likeList, err = db.GetVideoLikeList(vid); err != nil {
				errChan <- err
			}
			wg.Done()
		}(&data)
		go func(data *videoSyncData) {
			if data.commentCount, err = db.GetVideoCommentCount(vid); err != nil {
				errChan <- err
			}
			wg.Done()
		}(&data)
		wg.Wait()
		select {
		case result := <-errChan:
			return result
		default:
		}
		syncList = append(syncList, data)
	}

	errChan = make(chan error, 2)
	wg.Add(2)
	go func(syncList *[]videoSyncData) {
		if err := videoSyncDB2Redis(syncList); err != nil {
			errChan <- err
		}
		wg.Done()
	}(&syncList)
	go func(syncList *[]videoSyncData) {
		if err := vidoeSyncDB2Elastic(syncList); err != nil {
			errChan <- err
		}
		wg.Done()
	}(&syncList)
	wg.Wait()
	select {
	case result := <-errChan:
		return result
	default:
	}
	return nil
}

func videoSyncDB2Redis(syncList *[]videoSyncData) error {
	for _, item := range *syncList {
		if err := redis.PutVideoLikeInfo(item.vid, item.likeList); err != nil {
			return err
		}
	}
	return nil
}

func vidoeSyncDB2Elastic(syncList *[]videoSyncData) error {
	for _, item := range *syncList {
		if err := elasticsearch.UpdateVideoCommentAndLikeCount(item.vid, fmt.Sprint(len(*item.likeList)), fmt.Sprint(item.commentCount)); err != nil {
			return err
		}
	}
	return nil
}
