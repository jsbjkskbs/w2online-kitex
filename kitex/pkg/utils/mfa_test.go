package utils_test

import (
	"testing"
	"work/pkg/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestGenerateTOTP(t *testing.T) {
	hlog.Info(utils.NewAuthController(`10001`, ``, ``).GenerateTOTP())
}

func TestVerify(t *testing.T) {
	code := ``
	secret := ``
	uid := ``
	hlog.Info(utils.NewAuthController(uid, code, secret).VerifyTOTP())
}
