package handler_interact

import (
	"context"
	"work/kitex_gen/interact"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_interact/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_interact "work/rpc/facade/model/base/interact"

	"github.com/cloudwego/hertz/pkg/app"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_interact.CommentListRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	data, err := client.CommentList(ctx, &interact.CommentListRequest{
		VideoId:   facadeReq.VideoId,
		CommentId: facadeReq.CommentId,
		PageSize:  facadeReq.PageSize,
		PageNum:   facadeReq.PageNum,
	})
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_interact.CommentListResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Data: &facade_interact.CommentListResponse_CommentListResponseData{
			Items: *convert.KitexGenToRespComment(&data.Items),
		},
	})
}
