package main

import (
	"douyin/cmd/core/dal"
	core "douyin/kitex_gen/core/coreservice"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func Init() {
	dal.Init()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}

	Init()

	svr := core.NewServer(new(CoreServiceImpl),
		server.WithServiceAddr(addr), // address
	)

	// db.CreateVideo(
	// 	context.Background(),
	// 	[]*db.Video{
	// 		{
	// 			AuthorID:      1,
	// 			PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
	// 			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	// 			FavoriteCount: 0,
	// 			CommentCount:  0,
	// 			IsFavorite:    false,
	// 			Title:         "first",
	// 		},
	// 	},
	// )

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}

	// db.CreateUser(
	// 	context.Background(),
	// 	[]*db.User{
	// 		{
	// 			Name:          "TestUser",
	// 			FollowCount:   0,
	// 			FollowerCount: 0,
	// 			IsFollow:      false,
	// 		},
	// 	},
	// )

	// fmt.Println(db.GetVideoById(context.Background(), 1))
}
