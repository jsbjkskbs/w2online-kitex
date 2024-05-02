package handler_video

import (
	"context"
	"work/kitex_gen/video"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_video/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_video "work/rpc/facade/model/base/video"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoPopular(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_video.VideoPopularRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	data, err := client.VideoPopular(ctx, &video.VideoPopularRequest{
		PageNum:  facadeReq.PageNum,
		PageSize: facadeReq.PageSize,
	})
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoPopularResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &facade_video.VideoPopularResponse_VideoPopularResponseData{
			Items: *convert.KitexGenToRespVideo(&data.Items),
		},
	})
}
