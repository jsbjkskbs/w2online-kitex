package handler_user

import (
	"context"
	"work/kitex_gen/user"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	facade_user "work/rpc/facade/model/base/user"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthMfaBind(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_user.AuthMfaBindRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	req := user.AuthMfaBindRequest{
		Code:   facadeReq.Code,
		Secret: facadeReq.Secret,
	}
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	err = client.AuthMfaBind(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendResponse(c, errno.NoError, nil)
}
