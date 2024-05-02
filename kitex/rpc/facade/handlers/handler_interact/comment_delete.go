package handler_interact

import (
	"context"
	"work/kitex_gen/interact"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_interact "work/rpc/facade/model/base/interact"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func CommentDelete(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_interact.CommentDeleteRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	uid, err := jwt.CovertJWTPayloadToString(ctx, c)
	if err != nil {
		handlers.SendResponse(c, errno.TokenInvailed, nil)
		return
	}

	err = client.CommentDelete(ctx, &interact.CommentDeleteRequest{
		VideoId:    facadeReq.VideoId,
		CommentId:  facadeReq.CommentId,
		FromUserId: uid,
	})
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_interact.CommentDeleteResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
	})
}
