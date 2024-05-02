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
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func AvatarUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserAvatarUploadRequest

	if req.UserId, err = jwt.CovertJWTPayloadToString(ctx, c); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	uploadRawData, err := c.FormFile("file")
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

	avatarRawData, err := ioutil.ReadAll(file)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	req.Data = avatarRawData
	req.Filesize = uploadRawData.Size

	data, err := client.UserAvatarUpload(ctx, &req)
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_user.UserAvatarUploadResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &facade_user.UserAvatarUploadResponse_UserAvatarUploadResponseData{
			Id:        data.Uid,
			Username:  data.Username,
			AvatarUrl: data.AvatarUrl,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			DeletedAt: data.DeletedAt,
		},
	})
}
