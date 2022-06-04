package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/core"
	"douyin/pkg/errno"
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
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				Code:    Err.ErrCode,
				Message: Err.ErrMsg,
			},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			Code:    errno.Success.ErrCode,
			Message: errno.Success.ErrMsg,
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
