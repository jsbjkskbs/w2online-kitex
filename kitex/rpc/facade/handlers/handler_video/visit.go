package handler_video

import (
	"context"
	"work/kitex_gen/video"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_video "work/rpc/facade/model/base/video"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoVisit(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_video.VideoVisitRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	var req video.VideoVisitRequest
	req.FromIp = c.ClientIP()
	data, err := client.VideoVisit(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoVisitResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Item: &base.Video{
			Id:           data.Id,
			UserId:       data.UserId,
			VideoUrl:     data.VideoUrl,
			CoverUrl:     data.CoverUrl,
			Title:        data.Title,
			Description:  data.Description,
			VisitCount:   data.VisitCount,
			LikeCount:    data.LikeCount,
			CommentCount: data.CommentCount,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
			DeletedAt:    data.DeletedAt,
		},
	})
}
