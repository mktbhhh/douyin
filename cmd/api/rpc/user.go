package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUserRpc() {
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithHostPorts("127.0.0.1:8890"),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// UserRegister create user info
func UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
