package auth

import (
	"context"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func Auth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		DoubleTokenAuthFunc(),
		//jwt.AccessTokenJwtMiddleware.MiddlewareFunc(),
	)
}

func DoubleTokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.IsAccessTokenAvailable(ctx, c) {
			if !jwt.IsRefreshTokenAvailable(ctx, c) {
				handlers.SendResponse(c, errno.TokenInvailed, nil)
				c.Abort()
				return
			}
			jwt.GenerateAccessToken(ctx, c)
		}

		c.Next(ctx)
	}
}
