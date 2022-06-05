package middleware

import (
	"context"
	"douyin/cmd/api/param"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/constants"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func InitJWT() {
	AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar param.UserLoginParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, param.Resp{
				Code:    int64(code),
				Message: message,
			})
		},
		//HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
		//	c.JSON(http.StatusUnauthorized)
		//},
	})
}
