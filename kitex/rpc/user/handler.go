package main

import (
	"context"
	"fmt"
	"work/kitex_gen/base"
	user "work/kitex_gen/user"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserRegisterResponse)

	_, err = service.NewUserService(ctx).NewRegisterEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserLoginResponse)
	resp.Data = &base.User{}

	data, err := service.NewUserService(ctx).NewLoginEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = &base.User{
		Uid:       fmt.Sprint(data.Uid),
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: fmt.Sprint(data.CreatedAt),
		DeletedAt: fmt.Sprint(data.DeletedAt),
		UpdatedAt: fmt.Sprint(data.UpdatedAt),
	}
	return resp, nil
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, request *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserInfoResponse)
	resp.Data = &base.User{}

	data, err := service.NewUserService(ctx).NewInfoEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = &base.User{
		Uid:       fmt.Sprint(data.Uid),
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: fmt.Sprint(data.CreatedAt),
		DeletedAt: fmt.Sprint(data.DeletedAt),
		UpdatedAt: fmt.Sprint(data.UpdatedAt),
	}
	return resp, nil
}

// AvatarUpload implements the UserServiceImpl interface.
func (s *UserServiceImpl) AvatarUpload(ctx context.Context, request *user.UserAvatarUploadRequest) (resp *user.UserAvatarUploadResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserAvatarUploadResponse)
	resp.Data = &base.User{}

	data, err := service.NewUserService(ctx).NewAvatarUploadEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = &base.User{
		Uid:       fmt.Sprint(data.Uid),
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: fmt.Sprint(data.CreatedAt),
		DeletedAt: fmt.Sprint(data.DeletedAt),
		UpdatedAt: fmt.Sprint(data.UpdatedAt),
	}
	return resp, nil
}

// AuthMfaQrcode implements the UserServiceImpl interface.
func (s *UserServiceImpl) AuthMfaQrcode(ctx context.Context, request *user.AuthMfaQrcodeRequest) (resp *user.AuthMfaQrcodeResponse, err error) {
	// TODO: Your code here...
	resp = new(user.AuthMfaQrcodeResponse)
	resp.Data = &user.Qrcode{}

	data, err := service.NewUserService(ctx).NewQrcodeEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data.Data
	return resp, nil
}

// AuthMfaBind implements the UserServiceImpl interface.
func (s *UserServiceImpl) AuthMfaBind(ctx context.Context, request *user.AuthMfaBindRequest) (resp *user.AuthMfaBindResponse, err error) {
	// TODO: Your code here...
	resp = new(user.AuthMfaBindResponse)

	err = service.NewUserService(ctx).NewMfaBindEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// UserImageSearch implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserImageSearch(ctx context.Context, request *user.UserImageSearchRequest) (resp *user.UserImageSearchResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserImageSearchResponse)

	url, err := service.NewUserService(ctx).NewImageSearchEvent(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = url
	return resp, nil
}
