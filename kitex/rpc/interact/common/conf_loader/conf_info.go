package conf_loader

var (
	configPath = `./rpc/interact/`
	configName = `config`
	configType = `yaml`
)

func ModifyConf(path, name, cType string) {
	configPath = path
	configName = name
	configType = cType
}
