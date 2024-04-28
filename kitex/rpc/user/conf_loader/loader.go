package conf_loader

import (
	"work/rpc/user/dal"
	"work/rpc/user/dal/db"
	"work/rpc/user/oss"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var globalConfig *viper.Viper

func Run() error {
	globalConfig = viper.New()
	globalConfig.SetConfigName("config")
	globalConfig.SetConfigType("yaml")
	globalConfig.AddConfigPath(configPath)
	if err := globalConfig.ReadInConfig(); err != nil {
		return err
	}
	loadConfig()

	go func() {
		globalConfig.WatchConfig()
		globalConfig.OnConfigChange(func(e fsnotify.Event) {
			hlog.Info("modifying the cfg file has detected")
			loadConfig()
		})
	}()

	return nil
}

func loadConfig() {
	db.Conf.MysqlDSN = globalConfig.GetString("MysqlDSN")
	hlog.Info("MysqlDSN:" + db.Conf.MysqlDSN)
	dal.Load()

	ossConfig := globalConfig.GetStringMapString("OSS")
	oss.Bucket = ossConfig["bucket"]
	oss.SecretKey = ossConfig["secretKey"]
	oss.AccessKey = ossConfig["accessKey"]
	oss.Url = ossConfig["url"]
	oss.DefaultAvatarUrl = globalConfig.GetString("DefaultAvatarUrl")
	oss.Load()
}
