package testenv

import (
	"work/rpc/video/common/conf_loader"
	"work/rpc/video/common/dustman"
)

func Init(confPath string) {
	conf_loader.ModifyConf(confPath, `config`, `yaml`)
	conf_loader.Init()
	dustman.NewFileDustman().Run()
	dustman.NewRedisDustman().Run()
}
