package db

import (
	"work/pkg/sharding"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	DB *gorm.DB
)

func Load() {
	var err error
	DB, err = gorm.Open(mysql.Open(Conf.MysqlDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
	if err = DB.Use(sharding.NewSharding("to_user_id", 4, "messages")); err != nil {
		panic(err)
	}

	hlog.Infof("Mysql connected successfully.")
}
