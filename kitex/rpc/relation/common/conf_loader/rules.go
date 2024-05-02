package conf_loader

import (
	"work/pkg/utils/conf_loader"
	"work/rpc/relation/dal"
	"work/rpc/relation/dal/db"

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
	}
}
