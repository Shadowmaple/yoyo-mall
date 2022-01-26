package model

import "time"

type AddressModel struct {
	ID         uint32
	UserID     uint32
	Name       string // 收货人姓名
	Tel        string
	Detail     string // 详情地址，暂时先这样
	IsDefault  bool   // 是否默认地址
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}
