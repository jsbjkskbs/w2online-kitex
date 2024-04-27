package main

import (
	"context"
	user "work/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, request *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AvatarUpload implements the UserServiceImpl interface.
func (s *UserServiceImpl) AvatarUpload(ctx context.Context, request *user.UserAvatarUploadRequest) (resp *user.UserAvatarUploadResponse, err error) {
	// TODO: Your code here...
	return
}

// AuthMfaQrcode implements the UserServiceImpl interface.
func (s *UserServiceImpl) AuthMfaQrcode(ctx context.Context, request *user.AuthMfaQrcodeRequest) (resp *user.AuthMfaQrcodeResponse, err error) {
	// TODO: Your code here...
	return
}

// AuthMfaBind implements the UserServiceImpl interface.
func (s *UserServiceImpl) AuthMfaBind(ctx context.Context, request *user.AuthMfaBindRequest) (resp *user.AuthMfaBindResponse, err error) {
	// TODO: Your code here...
	return
}
