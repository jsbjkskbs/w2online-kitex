package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var (
	redisDBCommentInfo, redisDBVideoInfo *redis.Client
)

func Load() {
	redisDBCommentInfo = redis.NewClient(&redis.Options{
		Addr:     CommentInfo.Addr,
		Password: CommentInfo.Pwd,
		DB:       CommentInfo.DB,
	})
	redisDBVideoInfo = redis.NewClient(&redis.Options{
		Addr:     VideoInfo.Addr,
		Password: VideoInfo.Pwd,
		DB:       VideoInfo.DB,
	})

	if _, err := redisDBCommentInfo.Ping().Result(); err != nil {
		panic(err)
	}
	if _, err := redisDBVideoInfo.Ping().Result(); err != nil {
		panic(err)
	}
	hlog.Info("Redis connected successfully.")
}
