package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"work/kitex_gen/user"
	"work/pkg/errmsg"
	"work/pkg/utils"
	"work/rpc/user/dal/db"
	"work/rpc/user/oss"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (service UserService) NewRegisterEvent(request *user.UserRegisterRequest) (uid string, err error) {
	exist, err := db.UserIsExistByUsername(request.Username)
	if err != nil {
		return ``, errmsg.ServiceError
	}
	if exist {
		return ``, errmsg.UsernameAlreadyExistError
	}
	uid, err = db.InsertUser(&db.User{
		Username:  request.Username,
		Password:  utils.EncryptBySHA256(request.Password),
		AvatarUrl: oss.DefaultAvatarUrl,
		CreatedAt: time.Now().Unix(),
		DeletedAt: 0,
		UpdatedAt: time.Now().Unix(),
		MfaEnable: false,
	})
	if err != nil {
		return ``, errmsg.ServiceError
	}
	return uid, nil
}

func (service UserService) NewLoginEvent(request *user.UserLoginRequest) (*db.User, error) {
	user, err := db.GetUserByUsernameAndPwd(request.Username, utils.EncryptBySHA256(request.Password))
	if err != nil {
		return nil, errmsg.ServiceError
	}
	if user.MfaEnable {
		passed := utils.NewAuthController(fmt.Sprint(user.Uid), request.Code, user.MfaSecret).VerifyTOTP()
		if !passed {
			return nil, errmsg.AuthenticatorError
		}
	}
	return user, nil
}

func (service UserService) NewInfoEvent(request *user.UserInfoRequest) (*db.User, error) {
	return db.GetUserByUid(request.UserId)
}

func (service UserService) NewAvatarUploadEvent(request *user.UserAvatarUploadRequest) (*db.User, error) {
	data, err := service.uploadAvatarToOss(request.UserId, request.Data, request.Filesize)
	if err != nil {
		return nil, errmsg.ServiceError
	}

	return data, nil
}

func (service UserService) NewQrcodeEvent(request *user.AuthMfaQrcodeRequest) (*user.AuthMfaQrcodeResponse, error) {
	authInfo, err := utils.NewAuthController(request.UserId, ``, ``).GenerateTOTP()
	if err != nil {
		return nil, errmsg.ServiceError
	}

	return &user.AuthMfaQrcodeResponse{
		Data: &user.Qrcode{
			Secret: authInfo.Secret,
			Qrcode: utils.EncodeUrlToBase64String(authInfo.Url),
		},
	}, nil
}

func (service UserService) NewMfaBindEvent(request *user.AuthMfaBindRequest) error {
	passed := utils.NewAuthController(request.UserId, request.Code, request.Secret).VerifyTOTP()
	if !passed {
		return errmsg.AuthenticatorError
	}

	if err := db.UpdateMfaSecret(request.UserId, request.Secret); err != nil {
		return errmsg.ServiceError
	}
	return nil
}

func (service UserService) uploadAvatarToOss(uid string, avatarRawData []byte, filesize int64) (*db.User, error) {
	fileType := http.DetectContentType(avatarRawData)
	switch fileType {
	case `image/png`, `image/jpg`, `image/jpeg`:
		{
			var avatarUrl string
			var err error
			if avatarUrl, err = oss.UploadAvatar(&avatarRawData, filesize, fmt.Sprint(uid), fileType); err != nil {
				return nil, errmsg.OssUploadError
			}
			data, err := db.UpdateAvatarUrl(uid, avatarUrl)
			if err != nil {
				return nil, errmsg.ServiceError
			}
			return data, nil
		}
	default:
		return nil, errmsg.FileFormatNotSupportError
	}
}
