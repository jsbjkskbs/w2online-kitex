package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestEncodeUrlToString(t *testing.T) {
	hlog.Info(utils.EncodeUrlToBase64String(`http://test.com`))
}
