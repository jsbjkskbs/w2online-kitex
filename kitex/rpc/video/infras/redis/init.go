package redis

import (
	"github.com/go-redis/redis"
)

var (
	redisDBVideoUpload *redis.Client
	redisDBVideoInfo   *redis.Client
)

func Load() {

	redisDBVideoUpload = redis.NewClient(&redis.Options{
		Addr:     VideoUpload.Addr,
		Password: VideoUpload.Pwd,
		DB:       VideoUpload.DB,
	})

	redisDBVideoInfo = redis.NewClient(&redis.Options{
		Addr:     VideoInfo.Addr,
		Password: VideoInfo.Pwd,
		DB:       VideoInfo.DB,
	})

	if _, err := redisDBVideoUpload.Ping().Result(); err != nil {
		panic(err)
	}
	if _, err := redisDBVideoInfo.Ping().Result(); err != nil {
		panic(err)
	}
}
