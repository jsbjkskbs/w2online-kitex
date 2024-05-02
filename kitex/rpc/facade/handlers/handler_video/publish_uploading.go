package handler_video

import (
	"context"
	"io/ioutil"
	"work/kitex_gen/video"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_video "work/rpc/facade/model/base/video"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoPublishUploading(ctx context.Context, c *app.RequestContext) {
	var err error
	var facadeReq facade_video.VideoPublishUploadingRequest
	if err := c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	var req video.VideoPublishUploadingRequest
	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	rawData, err := c.FormFile("data")
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	file, err := rawData.Open()
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	req.ChunkNumber = facadeReq.ChunkNumber
	req.Data = data
	req.Filename = facadeReq.Filename
	req.IsM3u8 = facadeReq.IsM3U8
	req.Md5 = facadeReq.Md5
	req.Uuid = facadeReq.Uuid

	err = client.VideoPublishUploading(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoPublishUploadingResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
	})
}
