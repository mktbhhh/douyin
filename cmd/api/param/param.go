package param

import (
	"douyin/kitex_gen/core"
)

type Resp struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg,omitempty"`
}

type FeedResp struct {
	Resp
	VideoList []*core.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

type UserRegisterParam struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserRegisterResp struct {
	Resp
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginParam struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type GinUser struct {
	Id            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

type UserInfoResp struct {
	Resp
	User *GinUser `json:"user"`
}
