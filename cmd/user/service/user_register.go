package service

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"errors"
	"fmt"
	"io"
)

const (
	ErrId    = -1
	ErrToken = "err"
)

type UserRegisterService struct {
	ctx context.Context
}

// NewUserRegisterService new CreateUserService
func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

// UserRegister create user info.
func (s *UserRegisterService) UserRegister(req *user.DouyinUserRegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}

func (s *UserRegisterService) GetId(userName string) (int64, error) {
	userByUserName, err := db.GetUserByUserName(s.ctx, userName)
	if err != nil {
		return ErrId, err
	}

	if userByUserName == nil {
		return ErrId, errors.New("db register error")
	}

	return int64(userByUserName.ID), nil
}
