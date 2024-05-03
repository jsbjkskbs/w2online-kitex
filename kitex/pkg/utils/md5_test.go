package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestGetFileMD5(t *testing.T) {
	input := `./md5_test.go`
	hlog.Info(utils.GetFileMD5(input))
}

func TestGetBytesMD5(t *testing.T) {
	input := []byte(`hello world`)
	hlog.Info(utils.GetBytesMD5(input))
}
