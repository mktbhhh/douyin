package main

import (
	"context"
	"douyin/cmd/core/service"
	"douyin/kitex_gen/core"
	"time"
)

// CoreServiceImpl implements the last service interface defined in the IDL.
type CoreServiceImpl struct{}

// DouyinFeed implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) DouyinFeed(ctx context.Context, req *core.DouyinFeedRequest) (resp *core.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	resp = new(core.DouyinFeedResponse)

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	video, err := service.NewDouyinFeedService(ctx).DouyinFeed(req)
	if err != nil {
		resp.StatusCode = 100
		resp.StatusMsg = "error"
		return resp, nil
	}
	resp.VideoList = []*core.Video{video}
	resp.NextTime = time.Now().Unix()

	return resp, nil
}
