package db

import (
	"context"
	"douyin/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorID      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

func (v *Video) TabelName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

func GetVideoById(ctx context.Context, videoId int64) (Video, error) {
	var video Video
	err := DB.WithContext(ctx).Where("id = ?", videoId).Find(&video).Error
	return video, err
}
