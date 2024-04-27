package relation

import (
	"work/rpc/facade/handlers/handler_relation"
	"work/rpc/facade/router/auth"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	root := h.Group("/")
	{
		_follower := root.Group("/follower", auth.Auth()...)
		{
			_list := _follower.Group("/list")
			_list.GET("/", handler_relation.FollowerList)
		}
	}
	{
		_following := root.Group("/following", auth.Auth()...)
		{
			_list0 := _following.Group("/list")
			_list0.GET("/", handler_relation.FollowingList)
		}
	}
	{
		_friend := root.Group("/friend", auth.Auth()...)
		{
			_list1 := _friend.Group("/list")
			_list1.GET("/", handler_relation.FriendList)
		}
	}
	{
		_relation := root.Group("/relation", auth.Auth()...)
		{
			_action := _relation.Group("/action")
			_action.POST("/", handler_relation.RelationAction)
		}
	}
}
