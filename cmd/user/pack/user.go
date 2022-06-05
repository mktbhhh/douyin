package pack

import (
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
)

func User(dbUser *db.User) *user.User {
	if dbUser == nil {
		return nil
	}
	return &user.User{
		Id:            int64(dbUser.ID),
		Name:          dbUser.UserName,
		FollowCount:   dbUser.FollowerCount,
		FollowerCount: dbUser.FollowerCount,
		IsFollow:      false,
	}
}
