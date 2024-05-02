package handler_interact

import (
	"context"
	"work/kitex_gen/interact"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_interact "work/rpc/facade/model/base/interact"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func CommentPublish(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_interact.CommentPublishRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	var req interact.CommentPublishRequest
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	req.VideoId = facadeReq.VideoId
	req.CommentId = facadeReq.CommentId
	req.Content = facadeReq.Content

	err = client.CommentPublish(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_interact.CommentPublishResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
	})
}
