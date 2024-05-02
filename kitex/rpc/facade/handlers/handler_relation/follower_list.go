package handler_relation

import (
	"context"
	"work/kitex_gen/relation"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_relation/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_relation "work/rpc/facade/model/base/relation"

	"github.com/cloudwego/hertz/pkg/app"
)

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_relation.FollowerListRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	data, err := client.FollowerList(ctx, &relation.FollowerListRequest{
		UserId:   facadeReq.UserId,
		PageNum:  facadeReq.PageNum,
		PageSize: facadeReq.PageSize,
	})
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_relation.FollowerListResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &facade_relation.FollowerListResponse_FollowerListResponseData{
			Items: *convert.KitexGenToRespUserLite(&data.Items),
			Total: data.Total,
		},
	})
}
