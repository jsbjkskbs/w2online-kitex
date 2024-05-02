package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"work/kitex_gen/user"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/facade/infras/client"
	facade_user "work/rpc/facade/model/base/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

var (
	AccessTokenJwtMiddleware  *jwt.HertzJWTMiddleware
	RefreshTokenJwtMiddleware *jwt.HertzJWTMiddleware

	AccessTokenExpireTime  = time.Hour * 1
	RefreshTokenExpireTime = time.Hour * 72

	AccessTokenIdentityKey  = "access_token_field"
	RefreshTokenIdentityKey = "refresh_token_field"
)

type PayloadIdentityData struct {
	Uid string
}

func AccessTokenJwtInit() {
	var err error
	AccessTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:                         []byte("access_token_key_123456"),
		TokenLookup:                 "query:token,header:Access-Token",
		Timeout:                     AccessTokenExpireTime,
		IdentityKey:                 AccessTokenIdentityKey,
		WithoutDefaultTokenHeadName: true,

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest facade_user.UserLoginRequest
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			user, err := client.UserLogin(context.Background(), &user.UserLoginRequest{
				Username: loginRequest.Username,
				Password: loginRequest.Password,
			})
			if err != nil {
				return nil, err
			}
			if user == nil {
				return nil, errno.InfomationNotExist
			}
			c.Set("user_id", user.Uid)
			return PayloadIdentityData{Uid: fmt.Sprint(user.Uid)}, nil
		},

		// data为Authenticator返回的interface{}
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(PayloadIdentityData); ok {
				return jwt.MapClaims{
					AccessTokenJwtMiddleware.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
			hlog.CtxInfof(ctx, "Login Successfully. IP: "+c.ClientIP())
			c.Set("Access-Token", message)
		},

		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return utils.CreateBaseHttpResponse(e).StatusMsg
		},
	})

	if err != nil {
		panic(err)
	}
	hlog.Infof("Access-Token Jwt Initialized Successfully")
}

func GenerateAccessToken(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get(RefreshTokenJwtMiddleware.IdentityKey)
	data := PayloadIdentityData{
		Uid: v.(*PayloadIdentityData).Uid,
	}
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	c.Header("New-Access-Token", tokenString)
}

func IsAccessTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := AccessTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(AccessTokenJwtMiddleware.IdentityKey, identity)
	}
	if !AccessTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}
	return true

}

func ExtractUserIdWhenAuthorized(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	data, exist := c.Get(AccessTokenJwtMiddleware.IdentityKey)
	if !exist {
		return nil, errno.InfomationNotExist
	}
	return data, nil
}

func CovertJWTPayloadToString(ctx context.Context, c *app.RequestContext) (string, error) {
	data, err := ExtractUserIdWhenAuthorized(ctx, c)
	if err != nil {
		return ``, err
	}
	return data.(map[string]interface{})["Uid"].(string), nil
}

func RefreshTokenJwtInit() {
	var err error
	RefreshTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:                         []byte("refresh_token_key_abcdef"),
		TokenLookup:                 "query:Refresh-Token,header:Refresh-Token",
		Timeout:                     RefreshTokenExpireTime,
		IdentityKey:                 RefreshTokenIdentityKey,
		WithoutDefaultTokenHeadName: true,

		// 只在LoginHandler触发
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			uid, exist := c.Get("user_id")
			if !exist {
				return nil, errno.InfomationNotExist
			}
			return PayloadIdentityData{Uid: fmt.Sprint(uid)}, nil
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &PayloadIdentityData{
				Uid: claims[RefreshTokenJwtMiddleware.IdentityKey].(string),
			}
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(PayloadIdentityData); ok {
				return jwt.MapClaims{
					RefreshTokenJwtMiddleware.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
			c.Set("Refresh-Token", message)
		},

		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			c.Set("user_id", data.(*PayloadIdentityData).Uid)
			return true
		},
	})
	if err != nil {
		panic(err)
	}
	hlog.Infof("Refresh-Token Jwt Initialized Successfully")
}

func IsRefreshTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := RefreshTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(RefreshTokenJwtMiddleware.IdentityKey, identity)
	}
	if !RefreshTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}
	return true
}

func GenerateRefreshToken(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get(AccessTokenJwtMiddleware.IdentityKey)
	data := PayloadIdentityData{
		Uid: v.(*PayloadIdentityData).Uid,
	}
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	c.Header("New-Refresh-Token", tokenString)
}
