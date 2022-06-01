package rpc

import (
	"context"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/core/coreservice"
	"douyin/pkg/constants"
	"time"

	"github.com/cloudwego/kitex/client"
)

var coreClient coreservice.Client

func initCoreRpc() {
	c, err := coreservice.NewClient(
		constants.CoreServiceName,
		client.WithHostPorts("127.0.0.1:8889"),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
	)
	if err != nil {
		panic(err)
	}
	coreClient = c
}

func DouyinFeed(ctx context.Context, req *core.DouyinFeedRequest) ([]*core.Video, error) {
	resp, err := coreClient.DouyinFeed(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.VideoList, nil
}
