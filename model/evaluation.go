package model

import (
	"errors"
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type EvaluationModel struct {
	ID         uint32
	UserID     uint32
	OrderID    uint32
	ProductID  uint32
	Content    string
	Score      int8   // 评分，1-5
	Level      int8   // 0好评，1一般，2差评
	IsAnoymous bool   // 是否匿名
	Pictures   string // 图片，分号分割
	CreateTime time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}

func (e *EvaluationModel) TableName() string {
	return "evaluation"
}

func (e *EvaluationModel) Create() error {
	e.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(e).Error
}

func (e *EvaluationModel) Save() error {
	return DB.Self.Save(e).Error
}

func BatchCreateEvaluations(records []*EvaluationModel) error {
	return DB.Self.Create(records).Error
}

func GetEvaluationByID(id uint32) (*EvaluationModel, error) {
	var m EvaluationModel
	err := DB.Self.First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errno.ErrRecordNotFound
	}
	return &m, err
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
		query = query.Where("product_id = ?", productID)
	}

	err := query.Limit(limit).Offset(offset).Find(&list).Error

	return list, err
}
