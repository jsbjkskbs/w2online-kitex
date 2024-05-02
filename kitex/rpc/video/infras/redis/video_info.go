package redis

import (
	"strconv"

	"github.com/go-redis/redis"
)

func PutVideoLikeInfo(vid string, uidList *[]string) error {
	pipe := redisDBVideoInfo.TxPipeline()
	pipe.Del(`l:` + vid)
	pipe.Del(`nl:` + vid)
	for _, item := range *uidList {
		pipe.SAdd(`l:`+vid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func PutVideoVisitInfo(vid, visitCount string) error {
	score, _ := strconv.ParseFloat(visitCount, 64)
	_, err := redisDBVideoInfo.ZAdd(`visit`, redis.Z{Score: score, Member: vid}).Result()
	if err != nil {
		return err
	}
	return nil
}

func IncrVideoVisitInfo(vid string) error {
	_, err := redisDBVideoInfo.ZIncrBy(`visit`, 1, vid).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetVideoVisitCount(vid string) (int64, error) {
	_, err := redisDBVideoInfo.ZRank(`visit`, vid).Result()
	if err != nil {
		return -1, err
	}
	s, err := redisDBVideoInfo.ZScore(`visit`, vid).Result()
	if err != nil {
		return -1, err
	}
	return int64(s), nil
}

func GetVideoPopularList(pageNum, pageSize int64) (*[]string, error) {
	list, err := redisDBVideoInfo.ZRevRange(`visit`, (pageNum-1)*pageSize, pageNum*pageSize-1).Result()
	if err != nil {
		return nil, err
	}
	return &list, err
}

func IsVideoExist(vid string) bool {
	_, err := redisDBVideoInfo.ZScore(`visit`, vid).Result()
	return err == nil
}

func DeleteVideo(vid string) error {
	videoPipe := redisDBVideoInfo.TxPipeline()

	videoPipe.Del(`l:` + vid)
	videoPipe.Del(`nl:` + vid)
	videoPipe.ZRem(`visit`, vid)

	if _, err := videoPipe.Exec(); err != nil {
		return err
	}
	return nil
}
