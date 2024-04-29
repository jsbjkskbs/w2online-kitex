package handler_interact

import (
	"context"
	"work/kitex_gen/interact"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_video/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_interact "work/rpc/facade/model/base/interact"

	"github.com/cloudwego/hertz/pkg/app"
)

func LikeList(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_interact.LikeListRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	data, err := client.LikeList(ctx, &interact.LikeListRequest{
		UserId:   facadeReq.UserId,
		PageSize: facadeReq.PageSize,
		PageNum:  facadeReq.PageNum,
	})
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_interact.LikeListResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Data: &facade_interact.LikeListResponse_LikeListResponseData{
			Items: *convert.KitexGenToRespVideo(&data.Items),
		},
	})
}
