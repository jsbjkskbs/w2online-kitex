package handler_relation

import (
	"context"
	"work/kitex_gen/relation"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_relation "work/rpc/facade/model/base/relation"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_relation.RelationActionRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	var req relation.RelationActionRequest
	if req.FromUserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	req.ActionType = facadeReq.ActionType
	req.ToUserId = facadeReq.ToUserId

	err = client.RelationAction(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_relation.RelationActionResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
	})
}
