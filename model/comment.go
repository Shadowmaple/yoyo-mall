package model

import (
	"time"
	"yoyo-mall/pkg/errno"
)

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

func (c *CommentModel) TableName() string {
	return "comment"
}

func (c *CommentModel) Create() error {
	return DB.Self.Create(c).Error
}

func (c *CommentModel) Save() error {
	return DB.Self.Save(c).Error
}

func GetComments(evaluationID uint32, limit, offset int) ([]*CommentModel, error) {
	list := make([]*CommentModel, 0)

	d := DB.Self.Where("is_deleted = 0").Where("evaluation_id = ?", evaluationID).
		Limit(limit).Offset(offset).Find(&list)

	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}

	return list, nil
}

func CountComment(evaluationID uint32) int {
	var count int
	DB.Self.Where("is_deleted = 0").Where("evaluation_id = ?", evaluationID).Count(&count)
	return count
}
