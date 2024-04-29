package convert

import (
	"work/kitex_gen/base"
	resp "work/rpc/facade/model/base"
)

func KitexGenToRespVideo(items *[]*base.Video) *[]*resp.Video {
	result := make([]*resp.Video, 0)
	for _, item := range *items {
		result = append(result, &resp.Video{
			Id:           item.Id,
			UserId:       item.UserId,
			VideoUrl:     item.VideoUrl,
			CoverUrl:     item.CoverUrl,
			Title:        item.Title,
			Description:  item.Description,
			VisitCount:   item.VisitCount,
			LikeCount:    item.LikeCount,
			CommentCount: item.CommentCount,
			CreatedAt:    item.CreatedAt,
			DeletedAt:    item.DeletedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}
	return &result
}
