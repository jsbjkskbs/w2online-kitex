package syncman

import (
	"context"
	"work/rpc/video/dal/db"
	"work/rpc/video/infras/redis"
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
}

func (sm VideoSyncman) Stop() {
	sm.cancle()
}

type videoSyncData struct {
	vid          string
	visitCount   string
}

func videoSyncMwWhenInit() error {
	var (
		list *[]string
		err  error
	)
	for i := 0; !(len(*list) < 1000); i++ {
		list, err = db.GetVideoIdList(1000, int64(i))
		if err != nil {
			return err
		}

		var (
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
		if err := videoSyncDB2Redis(&syncList); err != nil {
			return err
		}
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
