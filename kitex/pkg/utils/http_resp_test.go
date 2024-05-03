package utils_test

import (
	"errors"
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestCreateHttpResponse(t *testing.T) {
	hlog.Info(utils.CreateBaseHttpResponse(errors.New(`test`)))
}
