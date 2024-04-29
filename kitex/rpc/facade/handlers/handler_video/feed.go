package handler_video

import (
	"context"
	"work/kitex_gen/video"
	"work/pkg/errmsg"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_video/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_video "work/rpc/facade/model/base/video"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoFeed(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_video.VideoFeedRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	data, err := client.VideoFeed(ctx, &video.VideoFeedRequest{
		LatestTime: facadeReq.LatestTime,
	})
	if err != nil {
		handlers.SendResponse(c, errmsg.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoFeedResponse{
		Base: &base.Status{
			Code: errmsg.NoError.ErrorCode,
			Msg:  errmsg.NoError.ErrorMsg,
		},
		Data: &facade_video.VideoFeedResponse_VideoFeedResponseData{
			Items: *convert.KitexGenToRespVideo(&data.Items),
		},
	})
}
