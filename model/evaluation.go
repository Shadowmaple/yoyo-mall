package model

import "time"

type EvaluationModel struct {
	ID         uint32
	UserID     uint32
	OrderID    uint32
	ProductID  uint32
	Content    string
	Score      int8   // 评分，1-5
	Rank       int8   // 0好评，1一般，2差评
	IsAnoymous bool   // 是否匿名
	Pictures   string // 图片，分号分割
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}
