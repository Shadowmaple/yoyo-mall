package model

import "time"

type UserModel struct {
	ID             uint32
	Nickname       string
	WechatUniqueID string // 微信唯一id
	Avatar         string
	Username       string
	Password       string
	State          int8 // 状态，0正常，1失效
	Role           int8 // 角色，0普通用户，1管理员，2商家
	CreateTime     *time.Time
}
