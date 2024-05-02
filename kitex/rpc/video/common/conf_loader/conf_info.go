package conf_loader

var (
	configPath = `./rpc/video/`
	configName = `config`
	configType = `yaml`
)

func ModifyConf(path, name, cType string) {
	configPath = path
	configName = name
	configType = cType
}
