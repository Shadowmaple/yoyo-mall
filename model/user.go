package model

import (
	"errors"
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type UserModel struct {
	ID             uint32
	Nickname       string
	WechatUniqueID string // 微信唯一id
	Avatar         string
	Gender         int8
	Username       string
	Password       string
	State          int8 // 状态，0正常，1失效
	Role           int8 // 角色，0普通用户，1管理员，2商家
	CreateTime     time.Time
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) Create() error {
	u.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(u).Error
}

func (u *UserModel) Save() error {
	return DB.Self.Save(u).Error
}

func GetUserByID(id uint32) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.First(u, "id = ?", id)
	if errors.Is(d.Error, gorm.ErrRecordNotFound) {
		return u, errno.ErrRecordNotFound
	}
	return u, nil
}

func GetUserByWechat(id string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.First(u, "wechat_unique_id = ?", id)
	if errors.Is(d.Error, gorm.ErrRecordNotFound) {
		return u, errno.ErrRecordNotFound
	}
	return u, d.Error
}
