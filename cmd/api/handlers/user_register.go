package handlers

import (
	"context"
	"douyin/cmd/api/middleware"
	"douyin/cmd/api/param"
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

func UserRegisterSeedResp(c *gin.Context, err error, user_id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, param.UserRegisterResp{
		Resp: param.Resp{
			Code:    Err.ErrCode,
			Message: Err.ErrMsg,
		},
		UserId: user_id,
		Token:  token,
	})
}

// UserRegister register user info
func UserRegister(c *gin.Context) {
	var registerVar param.UserRegisterParam
	if err := c.ShouldBind(&registerVar); err != nil {
		UserRegisterSeedResp(c, err, ErrUserId, ErrToken)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		UserRegisterSeedResp(c, errno.ParamErr, ErrUserId, ErrToken)
		return
	}

	userId, err := rpc.UserRegister(context.Background(), &user.DouyinUserRegisterRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		UserRegisterSeedResp(c, err, ErrUserId, ErrToken)
		return
	}

	mw := middleware.AuthMiddleware
	tokenString, _, err := mw.TokenGenerator(userId)
	if err != nil {
		UserRegisterSeedResp(c, err, userId, ErrToken)
		return
	}

	UserRegisterSeedResp(c, errno.Success, userId, tokenString)
}
