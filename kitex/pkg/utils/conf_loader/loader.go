package conf_loader

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	GlobalConfig *viper.Viper
	RuleTable    []*ConfLoadRule
)

type ConfLoader struct {
	RegisterParam []interface{}
	Register      func(v ...any)

	ConfigName string
	ConfigType string
	ConfigPath string
}

func NewConfLoader(register func(v ...any), param []interface{}, configName, configType, configPath string) *ConfLoader {
	return &ConfLoader{
		Register:      register,
		RegisterParam: param,
		ConfigName:    configName,
		ConfigType:    configType,
		ConfigPath:    configPath,
	}
}

func (c *ConfLoader) Run() error {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigName(c.ConfigName)
	GlobalConfig.SetConfigType(c.ConfigType)
	GlobalConfig.AddConfigPath(c.ConfigPath)
	if err := GlobalConfig.ReadInConfig(); err != nil {
		return err
	}

	c.Register(c.RegisterParam...)

	loadConfig()

	go func() {
		GlobalConfig.WatchConfig()
		GlobalConfig.OnConfigChange(func(e fsnotify.Event) {
			hlog.Info("The config has changed")
			loadConfig()
		})
	}()

	return nil
}

func loadConfig() {
	for _, rule := range RuleTable {
		if err := rule.ConfLoadMethod(rule.ConfLoadMethodParam...); err != nil {
			alarmer(rule.Level, rule.RuleName)
			rule.FailedTrigger(rule.FailedTriggerParam...)
		}
		rule.SuccessTrigger(rule.SuccessTriggerParam...)
	}
}

func alarmer(level int, rulename string) {
	switch level {
	case LevelError:
		hlog.Error(rulename + " failed to load.")
	case LevelWarn:
		hlog.Warn(rulename + " failed to load.")
	case LevelFatal:
		hlog.Fatal(rulename + " failed to load.")
	}
}
