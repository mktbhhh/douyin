package main

import (
	"douyin/cmd/user/dal"
	user "douyin/kitex_gen/user/userservice"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func Init() {
	dal.Init()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		panic(err)
	}

	Init()

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
