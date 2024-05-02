package elasticsearch_test

import (
	"testing"
	"work/rpc/video/infras/elasticsearch"
	testenv "work/rpc/video/testEnv"
)

var confPath = `../../`

func TestCreateDoc(t *testing.T){
	testenv.Init(confPath)
	elasticsearch.CreateVideoDoc(&elasticsearch.Video{
		Title: `标题`,
		Description: `描述`,
		Username: `cyk`,
		UserId: `10003`,
		CreatedAt: 1706269206,
		Info: elasticsearch.VideoOtherdata{
			Id: `10004`,
			VideoUrl: `s7rh811kd.hn-bkt.clouddn.com/video/10004/video.mp4`,
			CoverUrl: `s7rh811kd.hn-bkt.clouddn.com/video/10004/cover.jpg`,
			VisitCount: 0,
			LikeCount: 1,
			CommentCount: 0,
			UpdatedAt: 1710869608,
			DeletedAt: 0,
		},
	})
}
