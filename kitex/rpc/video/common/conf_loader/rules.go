package conf_loader

import (
	"work/pkg/utils/conf_loader"
	"work/rpc/video/dal"
	"work/rpc/video/dal/db"
	"work/rpc/video/infras/elasticsearch"
	"work/rpc/video/infras/redis"
	"work/rpc/video/infras/oss"
	"work/rpc/video/service"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func register(_ ...any) {
	conf_loader.RuleTable = []*conf_loader.ConfLoadRule{
		{
			RuleName: `MysqlConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				db.Conf.MysqlDSN = conf_loader.GlobalConfig.GetString("MysqlDSN")
				dal.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{
				&db.Conf.MysqlDSN,
			},
			SuccessTrigger: func(v ...any) {
				hlog.Info("MysqlDSN:" + *(v[0].(*string)))
			},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: `OssConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				ossConfig := conf_loader.GlobalConfig.GetStringMapString("OSS")
				oss.Bucket = ossConfig["bucket"]
				oss.SecretKey = ossConfig["secretKey"]
				oss.AccessKey = ossConfig["accessKey"]
				oss.Url = ossConfig["url"]
				oss.DefaultAvatarUrl = conf_loader.GlobalConfig.GetString("DefaultAvatarUrl")
				oss.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger: func(v ...any) {
				hlog.Info("oss config loaded successfully")
			},
		},
		{
			RuleName: `RedisConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				redisConfig := conf_loader.GlobalConfig.GetStringMap("redis")

				videoInfoConfig := redisConfig["video_info"].(map[string]interface{})
				videoUploadConfig := redisConfig["video_upload"].(map[string]interface{})

				redis.VideoInfo.Addr = videoInfoConfig["address"].(string)
				redis.VideoInfo.Pwd = videoInfoConfig["password"].(string)
				redis.VideoInfo.DB = videoInfoConfig["db"].(int)

				redis.VideoUpload.Addr = videoUploadConfig["address"].(string)
				redis.VideoUpload.Pwd = videoUploadConfig["password"].(string)
				redis.VideoUpload.DB = videoUploadConfig["db"].(int)

				redis.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger: func(v ...any) {
				hlog.Info("Redis connected successfully.")
			},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: `ElasticSearchConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				elasticsearch.ElasticAddr = conf_loader.GlobalConfig.GetString("ElasticAddr")
				elasticsearch.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger: func(v ...any) {
				hlog.Info("ElasticAddr:" + elasticsearch.ElasticAddr)
			},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: `TempVideoConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				service.TempVideoFolderPath = conf_loader.GlobalConfig.GetString("TempVideoFolderPath")
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger:      func(v ...any) {},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
	}
}
