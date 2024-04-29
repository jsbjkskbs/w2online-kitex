package handler_video

import (
	"context"
	"work/kitex_gen/video"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_video "work/rpc/facade/model/base/video"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoPublishStart(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_video.VideoPublishStartRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	var req video.VideoPublishStartRequest
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}
	req.Title = facadeReq.Title
	req.Description = facadeReq.Description
	req.ChunkTotalNumber = facadeReq.ChunkTotalNumber

	data, err := client.VideoPublishStart(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoPublishStartResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Uuid: data,
	})
}
