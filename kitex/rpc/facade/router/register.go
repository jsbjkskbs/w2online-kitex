package router

import (
	"work/rpc/facade/router/interact"
	"work/rpc/facade/router/relation"
	"work/rpc/facade/router/user"
	"work/rpc/facade/router/video"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	user.Register(h)

	video.Register(h)

	interact.Register(h)

	relation.Register(h)
}
