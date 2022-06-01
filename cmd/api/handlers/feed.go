package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/core"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []*core.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func DouyinFeed(c *gin.Context) {
	req := &core.DouyinFeedRequest{}

	videoList, err := rpc.DouyinFeed(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 100,
				StatusMsg:  "feed失败",
			},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
