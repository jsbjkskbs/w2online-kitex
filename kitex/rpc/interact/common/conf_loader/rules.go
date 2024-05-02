package conf_loader

import (
	"work/pkg/utils/conf_loader"
	"work/rpc/interact/dal"
	"work/rpc/interact/dal/db"
	"work/rpc/interact/infras/elasticsearch"
	"work/rpc/interact/infras/rabbitmq"
	"work/rpc/interact/infras/redis"

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
			RuleName: `RabbitMQConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				rabbitmq.RabbitmqDSN = conf_loader.GlobalConfig.GetString("RabbitmqDSN")
				rabbitmq.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger: func(v ...any) {
				hlog.Info("RabbitmqDSN:" + rabbitmq.RabbitmqDSN)
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
			RuleName: `RedisConfig`,
			Level:    conf_loader.LevelIgnore,

			ConfLoadMethodParam: []interface{}{},
			ConfLoadMethod: func(v ...any) error {
				redisConfig := conf_loader.GlobalConfig.GetStringMap("redis")

				commentInfoConfig := redisConfig["comment_info"].(map[string]interface{})

				redis.CommentInfo.Addr = commentInfoConfig["address"].(string)
				redis.CommentInfo.Pwd = commentInfoConfig["password"].(string)
				redis.CommentInfo.DB = commentInfoConfig["db"].(int)

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
	}
}
