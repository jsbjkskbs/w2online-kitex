package handler_wschat

import (
	"context"
	"work/kitex_gen/message"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"

	"github.com/cloudwego/hertz/pkg/app"
)

func InsertMessage(ctx context.Context, c *app.RequestContext) {
	var req message.InsertMessageRequest
	if err := c.BindAndValidate(&req); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	err := client.InsertMessage(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
	}

	handlers.SendResponse(c, errmsg.NoError, nil)
}
