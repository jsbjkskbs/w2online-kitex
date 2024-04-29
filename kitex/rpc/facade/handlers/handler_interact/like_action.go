package handler_interact

import (
	"context"
	"work/kitex_gen/interact"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_interact "work/rpc/facade/model/base/interact"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_interact.LikeActionRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	var req interact.LikeActionRequest
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}
	req.ActionType = facadeReq.ActionType
	req.CommentId = facadeReq.CommentId
	req.VideoId = facadeReq.VideoId

	err = client.LikeAction(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_interact.LikeActionResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
	})
}
