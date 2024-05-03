package convert_test

import (
	"testing"
	"work/kitex_gen/base"
	"work/rpc/facade/handlers/handler_relation/convert"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestConvert(t *testing.T) {
	list := make([]*base.UserLite, 0)
	list = append(list, &base.UserLite{
		Uid:       `1`,
		Username:  `2`,
		AvatarUrl: `3`,
	})
	hlog.Info(convert.KitexGenToRespUserLite(&list))
}
