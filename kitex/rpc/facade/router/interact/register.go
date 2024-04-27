package interact

import (
	"work/rpc/facade/handlers/handler_interact"
	"work/rpc/facade/router/auth"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	root := h.Group("/")
	{
		_comment := root.Group("/comment", auth.Auth()...)
		{
			_delete := _comment.Group("/delete")
			_delete.DELETE("/", handler_interact.CommentDelete)
		}
		{
			_list := _comment.Group("/list")
			_list.GET("/", handler_interact.CommentList)
		}
		{
			_publish := _comment.Group("/publish")
			_publish.POST("/", handler_interact.CommentPublish)
		}
	}
	{
		_like := root.Group("/like", auth.Auth()...)
		{
			_action := _like.Group("/action")
			_action.POST("/", handler_interact.LikeAction)
		}
		{
			_list0 := _like.Group("/list")
			_list0.GET("/", handler_interact.LikeList)
		}
	}
}
