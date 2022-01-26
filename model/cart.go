package model

import "time"

// 购物车
type CartModel struct {
	ID         uint32
	UserID     uint32
	ProductID  uint32
	Num        int // 商品数量
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}
