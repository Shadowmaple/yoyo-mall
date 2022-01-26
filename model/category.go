package model

import "time"

type CategoryModel struct {
	ID         uint32
	ParentID   uint32 // 父类目id，0则为根类目
	Name       string
	Order      int    // 排列顺序，0最大
	Image      string // 类目图片
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}
