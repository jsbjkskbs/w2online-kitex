package handler_user

import (
	"context"
	"work/kitex_gen/user"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_user "work/rpc/facade/model/base/user"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_user.UserRegisterRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	req := user.UserRegisterRequest{
		Username: facadeReq.Username,
		Password: facadeReq.Password,
	}
	err := client.UserRegister(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	jwt.AccessTokenJwtMiddleware.LoginHandler(ctx, c)
	jwt.RefreshTokenJwtMiddleware.LoginHandler(ctx, c)

	handlers.SendFormedResponse(c, &facade_user.UserRegisterResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		AccessToken:  c.GetString("Access-Token"),
		RefreshToken: c.GetString("Refresh-Token"),
	})
}
