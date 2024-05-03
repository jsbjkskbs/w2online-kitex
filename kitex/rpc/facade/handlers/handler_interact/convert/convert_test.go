package convert_test

import (
	"testing"
	"work/kitex_gen/base"
	"work/rpc/facade/handlers/handler_interact/convert"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TestConvert(t *testing.T) {
	list := make([]*base.Comment, 0)
	list = append(list, &base.Comment{
		Id:         `1`,
		UserId:     `2`,
		VideoId:    `3`,
		ParentId:   `4`,
		LikeCount:  5,
		ChildCount: 6,
		Content:    `7`,
		CreatedAt:  `8`,
		UpdatedAt:  `9`,
		DeletedAt:  `10`,
	})
	hlog.Info(convert.KitexGenToRespComment(&list))
}
