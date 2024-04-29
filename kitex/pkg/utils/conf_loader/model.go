package conf_loader

const (
	LevelIgnore = iota
	LevelWarn
	LevelError
	LevelFatal
)

type ConfLoadRule struct {
	RuleName string

	// 加载失败后应该看作? (LevelWarn,LevelError,LevelFatal etc.)
	Level int

	// 加载方法
	ConfLoadMethodParam []interface{}
	ConfLoadMethod      func(v ...any) error

	// 加载成功触发
	SuccessTriggerParam []interface{}
	SuccessTrigger      func(v ...any)

	// 加载失败触发
	FailedTriggerParam []interface{}
	FailedTrigger      func(v ...any)
}
