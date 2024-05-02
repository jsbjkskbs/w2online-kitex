package conf_loader

import "work/pkg/utils/conf_loader"

func Init() {
	err := conf_loader.NewConfLoader(
		register,
		[]interface{}{},
		configName,
		configType,
		configPath,
	).Run()
	if err != nil {
		panic(err)
	}
}
