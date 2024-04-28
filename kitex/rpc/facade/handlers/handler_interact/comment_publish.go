package handler_interact

import (
	"context"
	"work/kitex_gen/interact"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func CommentPublish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.CommentPublishRequest
	if err := c.BindAndValidate(&req); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	err = client.CommentPublish(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
	}

	handlers.SendResponse(c, errmsg.NoError, nil)
}
