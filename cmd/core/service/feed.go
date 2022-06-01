package service

import (
	"context"
	"douyin/cmd/core/dal/db"
	"douyin/cmd/core/pack"
	"douyin/kitex_gen/core"
)

type DouyinFeedService struct {
	ctx context.Context
}

func NewDouyinFeedService(ctx context.Context) *DouyinFeedService {
	return &DouyinFeedService{ctx: ctx}
}

func (s *DouyinFeedService) DouyinFeed(req *core.DouyinFeedRequest) (*core.Video, error) {
	modelVideos, err := db.GetVideoById(s.ctx, 1)
	if err != nil {
		return nil, err
	}
	return pack.Video(&modelVideos), nil
}
