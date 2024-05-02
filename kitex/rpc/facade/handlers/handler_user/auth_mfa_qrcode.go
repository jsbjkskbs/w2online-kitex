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

func AuthMfaQrcode(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_user.AuthMfaQrcodeRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	req := user.AuthMfaQrcodeRequest{}
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	data, err := client.AuthMfaQrcode(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_user.AuthMfaQrcodeResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &facade_user.AuthMfaQrcodeResponse_Qrcode{
			Secret: data.Secret,
			Qrcode: data.Qrcode,
		},
	})
}
