package conf_loader

import (
	"work/pkg/utils/conf_loader"
	"work/rpc/user/dal"
	"work/rpc/user/dal/db"
	"work/rpc/user/infras/oss"

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
	}
}
