package db

type Video struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	VideoUrl    string `json:"video_url"`
	CoverUrl    string `json:"cover_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VisitCount  int64  `json:"visit_count"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

func GetVideoIdList() (*[]string, error) {
	list := make([]string, 0)
	if err := DB.Table(`videos`).Select("id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}
