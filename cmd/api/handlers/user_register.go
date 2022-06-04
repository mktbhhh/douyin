package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrUserId = -10086
	ErrToken  = "err"
)

func UserRegisterSeedResponse(c *gin.Context, err error, user_id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: Response{
			Code:    Err.ErrCode,
			Message: Err.ErrMsg,
		},
		user_id: user_id,
		token:   token,
	})
}

// UserRegister register user info
func UserRegister(c *gin.Context) {

	var registerVar UserRegisterParam
	if err := c.ShouldBind(&registerVar); err != nil {
		UserRegisterSeedResponse(c, errno.ConvertErr(err), ErrUserId, ErrToken)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		UserRegisterSeedResponse(c, errno.ParamErr, ErrUserId, ErrToken)
		return
	}

	resp, err := rpc.UserRegister(context.Background(), &user.DouyinUserRegisterRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		UserRegisterSeedResponse(c, errno.ConvertErr(err), ErrUserId, ErrToken)
		return
	}

	UserRegisterSeedResponse(c, errno.Success, resp.UserId, resp.Token)
}
