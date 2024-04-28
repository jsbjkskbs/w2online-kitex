package handler_user

import (
	"context"
	"work/kitex_gen/user"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_user "work/rpc/facade/model/base/user"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_user.UserInfoRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	req := user.UserInfoRequest{
		UserId: facadeReq.UserId,
	}
	data, err := client.UserInfo(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_user.UserInfoResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Data: &base.User{
			Uid:       data.Uid,
			Username:  data.Username,
			AvatarUrl: data.AvatarUrl,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			DeletedAt: data.DeletedAt,
		},
	})
}
