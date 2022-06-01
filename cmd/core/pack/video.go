package pack

import (
	"douyin/cmd/core/dal/db"
	"douyin/kitex_gen/core"
)

func Video(mvideo *db.Video) *core.Video {
	if mvideo == nil {
		return nil
	}
	return &core.Video{
		Id:            int64(mvideo.Model.ID),
		AuthorId:      mvideo.AuthorID,
		PlayUrl:       mvideo.PlayUrl,
		CoverUrl:      mvideo.CoverUrl,
		FavoriteCount: mvideo.FavoriteCount,
		CommentCount:  mvideo.CommentCount,
		IsFavorite:    mvideo.IsFavorite,
		Title:         mvideo.Title,
	}
}
