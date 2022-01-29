package model

import (
	"time"
	"yoyo-mall/pkg/errno"
)

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
	CreateTime time.Time
	IsDeleted  bool
	DeleteTime time.Time
}

func (e *EvaluationModel) TableName() string {
	return "evaluation"
}

func (e *EvaluationModel) Create() error {
	return DB.Self.Create(e).Error
}

func (e *EvaluationModel) Save() error {
	return DB.Self.Save(e).Error
}

func GetEvaluationList(userID, orderID, productID uint32, limit, offset int) ([]*EvaluationModel, error) {
	list := make([]*EvaluationModel, 0)
	query := DB.Self.Where("is_deleted = 0")
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if orderID > 0 {
		query = query.Where("order_id = ?", orderID)
	}
	if productID > 0 {
		query = query.Where("query_id = ?", productID)
	}

	d := query.Limit(limit).Offset(offset).Find(&list)
	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}

	return list, nil
}
