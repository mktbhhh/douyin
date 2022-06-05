package handlers

import (
	"context"
	"douyin/cmd/api/param"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/unknwon/com"

func UserInfoSeedResp(c *gin.Context, err error, userById *param.GinUser) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, param.UserInfoResp{
		Resp: param.Resp{
			Code:    Err.ErrCode,
			Message: Err.ErrMsg,
		},
		User: userById,
	})
}

func GetUserInfoById(c *gin.Context) {
	userId := com.StrTo(c.Query("user_id")).MustInt64()

	userById, err := rpc.UserInfo(
		context.Background(),
		&user.DouyinUserRequest{
			UserId: userId,
		})
	if err != nil {
		UserInfoSeedResp(c, err, nil)
	}

	UserInfoSeedResp(c, errno.Success, &param.GinUser{
		Id:            userById.Id,
		Name:          userById.Name,
		FollowCount:   userById.FollowCount,
		FollowerCount: userById.FollowerCount,
		IsFollow:      userById.IsFollow,
	})
}
