package redis

import (
	"fmt"
	"strconv"
	"time"
	"work/pkg/errno"

	"github.com/go-redis/redis"
)

func GetVideoDBKeys() ([]string, error) {
	keys, err := redisDBVideoUpload.Keys(`*`).Result()
	if err != nil {
		return nil, err
	}
	return keys, err
}

func DelVideoDBKeys(keys []string) error {
	pipe := redisDBVideoUpload.TxPipeline()
	for _, key := range keys {
		pipe.Del(key)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func NewVideoEvent(title, description, uid, chuckTotalNumber string) (string, error) {
	uuid := fmt.Sprint(time.Now().Unix())
	exist, err := redisDBVideoUpload.Exists("l:" + uid + ":" + uuid).Result()
	if err != nil {
		return ``, err
	}
	if exist != 0 {
		return ``, errno.InfomationAlreadyExistError
	}
	if _, err := redisDBVideoUpload.RPush("l:"+uid+":"+uuid, chuckTotalNumber, title, description).Result(); err != nil {
		return ``, err
	}
	return uuid, nil
}

func DoneChunkEvent(uuid, uid string, chunk int64) error {
	bitrecord, err := redisDBVideoUpload.GetBit("b:"+uid+":"+uuid, chunk).Result()
	if err != nil {
		return err
	}
	if bitrecord == 1 {
		return errno.InfomationAlreadyExistError
	}
	if _, err = redisDBVideoUpload.SetBit("b:"+uid+":"+uuid, chunk, 1).Result(); err != nil {
		return err
	}
	return nil
}

func IsChunkAllRecorded(uuid, uid string) (bool, error) {
	r, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 0, 0).Result()
	if err != nil {
		return false, err
	}
	chunkTotalNumber, _ := strconv.ParseInt(r[0], 10, 64)
	recordNumber, err := redisDBVideoUpload.BitCount("b:"+uid+":"+uuid, &redis.BitCount{
		Start: 0,
		End:   chunkTotalNumber - 1,
	}).Result()
	if err != nil {
		return false, err
	}
	return chunkTotalNumber == recordNumber, nil
}

func RecordM3U8Filename(uuid, uid, filename string) error {
	exist, err := redisDBVideoUpload.Exists("l:" + uid + ":" + uuid).Result()
	if err != nil {
		return err
	}
	if exist == 0 {
		return errno.RequestError
	}
	fLen, err := redisDBVideoUpload.LLen("l:" + uid + ":" + uuid).Result()
	if err != nil {
		return err
	}
	if fLen == 4 {
		return errno.RequestError
	}
	if _, err := redisDBVideoUpload.RPush("l:"+uid+":"+uuid, filename).Result(); err != nil {
		return err
	}
	return nil
}

func GetM3U8Filename(uuid, uid string) (string, error) {
	if filename, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 3, 3).Result(); err != nil || filename[0] == `` {
		return ``, errno.RequestError
	} else {
		return filename[0], nil
	}
}

func FinishVideoEvent(uuid, uid string) ([]string, error) {
	info, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 1, 2).Result()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func DeleteVideoEvent(uuid, uid string) error {
	pipe := redisDBVideoUpload.TxPipeline()
	pipe.Del("l:" + uid + ":" + uuid)
	pipe.Del("b:" + uid + ":" + uuid)
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}
