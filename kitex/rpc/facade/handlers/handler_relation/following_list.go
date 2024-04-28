package handler_relation

import (
	"context"
	"work/kitex_gen/relation"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"

	"github.com/cloudwego/hertz/pkg/app"
)

func FollowingList(ctx context.Context, c *app.RequestContext) {
	var req relation.FollowingListRequest
	if err := c.BindAndValidate(&req); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	data, err := client.FollowingList(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
	}

	handlers.SendResponse(c, errmsg.NoError, &map[string]interface{}{
		"data": data,
	})
}
