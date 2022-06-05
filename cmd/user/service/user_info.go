package service

import (
	"context"
	"douyin/cmd/user/dal/db"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new CreateUserService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) GetUserInfoById(userId int64) (*db.User, error) {
	userById, err := db.GetUserById(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	return userById, nil
}
