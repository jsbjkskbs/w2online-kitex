package handler_user

import (
	"context"
	"work/kitex_gen/user"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_user "work/rpc/facade/model/base/user"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_user.UserLoginRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	req := user.UserLoginRequest{
		Username: facadeReq.Username,
		Password: facadeReq.Password,
		Code:     facadeReq.Code,
	}
	data, err := client.UserLogin(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	jwt.AccessTokenJwtMiddleware.LoginHandler(ctx, c)
	jwt.RefreshTokenJwtMiddleware.LoginHandler(ctx, c)

	handlers.SendFormedResponse(c, &facade_user.UserLoginResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &base.User{
			Uid:       data.Uid,
			Username:  data.Username,
			AvatarUrl: data.AvatarUrl,
			CreatedAt: data.CreatedAt,
			DeletedAt: data.DeletedAt,
			UpdatedAt: data.UpdatedAt,
		},
		AccessToken:  c.GetString("Access-Token"),
		RefreshToken: c.GetString("Refresh-Token"),
	})
}
