package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestM3u8ToMp4(t *testing.T) {
	input := ``
	output := ``
	hlog.Info(utils.M3u8ToMp4(input, output))
}

func TestGenerateMp4CoverJpg(t *testing.T) {
	input := ``
	output := ``
	hlog.Info(utils.GenerateMp4CoverJpg(input, output))
}
