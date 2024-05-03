package conf_loader_test

import (
	"testing"
	"time"
	"work/pkg/utils/conf_loader"
)

func TestNewLoader(t *testing.T) {
	conf_loader.NewConfLoader(conf_loader.Register, []interface{}{}, `test`, `yaml`, `.`).Run()
	time.Sleep(2 * time.Second)
}
