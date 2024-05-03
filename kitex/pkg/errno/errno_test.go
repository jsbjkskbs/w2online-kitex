package errno_test

import (
	"errors"
	"testing"
	"work/pkg/errno"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestErrnoError(t *testing.T) {
	e := errno.Errno{Code: 1, Message: `test`}
	hlog.Info(e.Error())
}

func TestErrnoWithMessage(t *testing.T) {
	e := errno.BaseError
	newE := e.WithMessage(`test`)
	hlog.Info(newE)
}

func TestNewErrorMessage(t *testing.T) {
	hlog.Info(errno.NewErrorMessage(1, `test`))
}

func TestConvert(t *testing.T) {
	hlog.Info(errno.Convert(errors.New(`test`)))
}