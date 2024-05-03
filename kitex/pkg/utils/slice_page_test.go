package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestSlicePage(t *testing.T) {
	hlog.Info(utils.SlicePage(0, 10, 20))
	hlog.Info(utils.SlicePage(0, 30, 20))
	hlog.Info(utils.SlicePage(3, 8, 20))
}
