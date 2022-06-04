package service

import "context"

type UserInfoService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}
