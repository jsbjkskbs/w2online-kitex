package convert_test

import (
	"testing"
	"work/kitex_gen/base"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestConvert(t *testing.T) {
	list := make([]*base.Video, 0)
	list = append(list, &base.Video{
		Id:        `123`,
		CreatedAt: `231`,
	})
	hlog.Info(list)
}
