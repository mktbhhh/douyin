package handlers

import (
	"context"
	"douyin/cmd/api/param"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/core"
	"douyin/pkg/errno"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DouyinFeedSeedResp(c *gin.Context, err error, videoList []*core.Video) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, param.FeedResp{
		Resp: param.Resp{
			Code:    Err.ErrCode,
			Message: Err.ErrMsg,
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}

// DouyinFeed same demo video list for every request
func DouyinFeed(c *gin.Context) {
	req := &core.DouyinFeedRequest{}

	videoList, err := rpc.DouyinFeed(context.Background(), req)
	if err != nil {
		DouyinFeedSeedResp(c, err, nil)
		return
	}

	DouyinFeedSeedResp(c, errno.Success, videoList)
}
