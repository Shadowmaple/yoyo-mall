package model

import "time"

type CommentModel struct {
	ID           uint32
	UserID       uint32
	EvaluationID uint32 // 评价id
	ParentID     uint32 // 父评论id，暂时不做空着
	ReplyUserID  uint32 // 向谁回复，暂时空着不做
	Content      string
	IsAnoymous   bool // 是否匿名
	CreateTime   *time.Time
	IsDeleted    bool
	DeleteTime   *time.Time
}
