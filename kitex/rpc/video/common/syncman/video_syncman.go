package syncman

import (
	"context"
	"fmt"
	"sync"
	"time"
	"work/rpc/video/dal/db"
	"work/rpc/video/infras/elasticsearch"
	"work/rpc/video/infras/redis"

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
				visitCount int64
				vidList    *[]string
			)
			var err error
			if vidList, err = db.GetVideoIdList(); err != nil {
				hlog.Warn(err)
			}
			for _, vid := range *vidList {
				var err error
				if visitCount, err = redis.GetVideoVisitCount(vid); err != nil {
					hlog.Warn(err)
				}
				err = elasticsearch.UpdateVideoVisitCount(vid, fmt.Sprint(visitCount))
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
	vid        string
	visitCount string
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
		if data.visitCount, err = db.GetVideoVisitCount(vid); err != nil {
			return err
		}
		syncList = append(syncList, data)
	}

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
		if err := redis.PutVideoVisitInfo(item.vid, item.visitCount); err != nil {
			return err
		}
	}
	return nil
}

func vidoeSyncDB2Elastic(syncList *[]videoSyncData) error {
	for _, item := range *syncList {
		if err := elasticsearch.UpdateVideoVisitCount(item.vid, item.visitCount); err != nil {
			return err
		}
	}
	return nil
}
