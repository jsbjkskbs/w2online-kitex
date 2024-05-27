package user

import (
	"work/rpc/facade/handlers/handler_user"
	"work/rpc/facade/router/auth"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	root := h.Group("/")
	{
		_auth := root.Group("/auth", auth.Auth()...)
		{
			_mfa := _auth.Group("/mfa")
			{
				_bind := _mfa.Group("/bind")
				_bind.POST("/", handler_user.AuthMfaBind)
			}
			{
				_qrcode := _mfa.Group("/qrcode")
				_qrcode.GET("/", handler_user.AuthMfaQrcode)
			}
		}
	}
	{
		_user := root.Group("/user")
		{
			_avatar := _user.Group("/avatar", auth.Auth()...)
			{
				_upload := _avatar.Group("/upload")
				_upload.PUT("/", handler_user.AvatarUpload)
			}
		}
		{
			_info := _user.Group("/info", auth.Auth()...)
			_info.GET("/", handler_user.UserInfo)
		}
		{
			_login := _user.Group("/login")
			_login.POST("/", handler_user.UserLogin)
		}
		{
			_register := _user.Group("/register")
			_register.POST("/", handler_user.UserRegister)
		}
		{
			_image := _user.Group("/image")
			{
				_search := _image.Group("/search")
				_search.PUT("/", handler_user.UserImageSearch)
			}
		}
	}
}
