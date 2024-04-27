package handler_user

import (
	"context"
	"work/kitex_gen/user"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req user.UserLoginRequest
	if err := c.BindAndValidate(&req); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	data, err := client.UserLogin(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
	}

	handlers.SendResponse(c, errmsg.NoError, map[string]interface{}{
		"data":          data,
		"Access-Token":  c.GetString("Access-Token"),
		"Refresh-Token": c.GetString("Refresh-Token"),
	})
}
