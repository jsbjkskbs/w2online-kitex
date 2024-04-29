package handler_relation

import (
	"context"
	"work/kitex_gen/relation"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_relation/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_relation "work/rpc/facade/model/base/relation"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_relation.FriendListRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	var req relation.FriendListRequest
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}
	req.PageNum = facadeReq.PageNum
	req.PageSize = facadeReq.PageSize

	data, err := client.FriendList(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_relation.FriendListResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Data: &facade_relation.FriendListResponse_FriendListResponseData{
			Items: *convert.KitexGenToRespUserLite(&data.Items),
			Total: data.Total,
		},
	})
}
