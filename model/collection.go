package model

import "time"

type ColletionModel struct {
	ID         uint32
	UserID     uint32
	ProductID  uint32
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}
