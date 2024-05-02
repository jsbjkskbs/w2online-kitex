package elasticsearch

import (
	"context"
	"strconv"

	"github.com/olivere/elastic/v7"
)

type VideoOtherdata struct {
	Id           string `json:"id"`
	VideoUrl     string `json:"video_url"`
	CoverUrl     string `json:"cover_url"`
	VisitCount   int64  `json:"visit_count"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	UpdatedAt    int64  `json:"updated_at"`
	DeletedAt    int64  `json:"deleted_at"`
}

type Video struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Username    string         `json:"username"`
	UserId      string         `json:"user_id"`
	CreatedAt   int64          `json:"created_at"`
	Info        VideoOtherdata `json:"info"`
}

func UpdateVideoCommentAndLikeCount(vid, likeCount, commentCount string) error {
	bulk := elasticClient.Bulk()
	var (
		newLikeCount, _    = strconv.ParseInt(likeCount, 10, 64)
		newCommentCount, _ = strconv.ParseInt(commentCount, 10, 64)
	)
	lRequest := elastic.NewBulkUpdateRequest().Index("video").Type("_doc").Id(vid).
		Script(elastic.NewScript(`"ctx._source.info.like_count=params.new_like_count"`).Param("new_like_count", newLikeCount))
	cRequest := elastic.NewBulkUpdateRequest().Index("video").Type("_doc").Id(vid).
		Script(elastic.NewScript(`"ctx._source.info.comment_count=params.new_comment_count"`).Param("new_comment_count", newCommentCount))
	bulk.Add(lRequest, cRequest)
	if _, err := bulk.Do(context.Background()); err != nil {
		return err
	}
	return nil
}
