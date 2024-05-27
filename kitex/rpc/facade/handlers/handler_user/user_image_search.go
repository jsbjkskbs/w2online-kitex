package handler_user

import (
	"context"
	"io/ioutil"
	"work/kitex_gen/user"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_user "work/rpc/facade/model/base/user"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserImageSearch(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserImageSearchRequest

	uploadRawData, err := c.FormFile("data")
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	file, err := uploadRawData.Open()
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	req.Data = rawData

	data, err := client.ImageSearch(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_user.UserImageSearchResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: data,
	})
}
