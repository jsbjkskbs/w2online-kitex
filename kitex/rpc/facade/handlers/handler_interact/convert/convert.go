package convert

import (
	"work/kitex_gen/base"
	resp "work/rpc/facade/model/base"
)

func KitexGenToRespComment(items *[]*base.Comment) *[]*resp.Comment {
	result := make([]*resp.Comment, 0)
	for _, item := range *items {
		result = append(result, &resp.Comment{
			Id:         item.Id,
			UserId:     item.UserId,
			VideoId:    item.VideoId,
			ParentId:   item.ParentId,
			Content:    item.Content,
			ChildCount: item.ChildCount,
			LikeCount:  item.LikeCount,
			CreatedAt:  item.CreatedAt,
			DeletedAt:  item.DeletedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}
	return &result
}
