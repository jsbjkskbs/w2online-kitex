package handler_video

import (
	"context"
	"work/kitex_gen/video"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoPublishComplete(ctx context.Context, c *app.RequestContext) {
	var req video.VideoPublishCompleteRequest
	if err := c.BindAndValidate(&req); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	err := client.VideoPublishComplete(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
	}

	handlers.SendResponse(c, errmsg.NoError, nil)
}