package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestConvertTimestampToStringDefault(t *testing.T) {
	hlog.Info(utils.ConvertTimestampToStringDefault(0))
}

func TestConvertTimestampToString(t *testing.T) {
	hlog.Info(utils.ConvertTimestampToString(0,`05:15:04 01-2006-02`))
}

func TestConvertStringToTimestampDefault(t *testing.T) {
	hlog.Info(utils.CovertStringToTimestampDefault(`1970-01-01 08:00:00`))
}