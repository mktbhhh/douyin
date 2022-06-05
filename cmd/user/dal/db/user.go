package db

import (
	"context"
	"douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `gorm:"type:varchar(32);not null" json:"user_name"`
	Avatar        string `gorm:"type:varchar(512);not null" json:"avatar"`
	Password      string `gorm:"type:varchar(512);not null" json:"password"`
	FollowCount   int64  `gorm:"type:int;DEFAULT:0" json:"follow_count"`
	FollowerCount int64  `gorm:"type:int;DEFAULT:0" json:"follower_count"`
	// IsFollow      bool   `json:"is_follow"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserByUserName(ctx context.Context, userName string) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserById(ctx context.Context, userId int64) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
