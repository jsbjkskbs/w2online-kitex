package testenv

import (
	"work/rpc/video/common/conf_loader"
	"work/rpc/video/common/dustman"
	"work/rpc/video/common/syncman"
)

func Init(confPath string) {
	conf_loader.ModifyConf(confPath, `config`, `yaml`)
	conf_loader.Init()
	syncman.NewVideoSyncman().Run()
	dustman.NewFileDustman().Run()
	dustman.NewRedisDustman().Run()
}
