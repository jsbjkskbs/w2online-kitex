package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestEncryptBySHA256(t *testing.T) {
	hlog.Info(utils.EncryptBySHA256(`123456789`))
}

// RSA testing is done by api test
