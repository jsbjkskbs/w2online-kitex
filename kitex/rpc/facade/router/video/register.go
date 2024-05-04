package video

import (
	"work/rpc/facade/handlers/handler_video"
	"work/rpc/facade/router/auth"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	root := h.Group("/")
	{
		_video := root.Group("/video")
		{
			_feed := _video.Group("/feed")
			_feed.GET("/", handler_video.VideoFeed)
		}
		{
			_list := _video.Group("/list", auth.Auth()...)
			_list.GET("/", handler_video.VideoList)
		}
		{
			_popular := _video.Group("/popular")
			_popular.GET("/", handler_video.VideoPopular)
		}
		{
			_publish := _video.Group("/publish", auth.Auth()...)
			_publish.POST("/cancle", handler_video.VideoPublishCancle)
			_publish.POST("/complete", handler_video.VideoPublishComplete)
			_publish.POST("/start", handler_video.VideoPublishStart)
			_publish.POST("/uploading", handler_video.VideoPublishUploading)
		}
		{
			_search := _video.Group("/search", auth.Auth()...)
			_search.POST("/", handler_video.VideoSearch)
		}
		{
			_visit := _video.Group("/visit")
			_visit.GET("/:id", handler_video.VideoVisit)
		}
	}
}
