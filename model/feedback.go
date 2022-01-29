package model

import (
	"time"
	"yoyo-mall/util"
)

type FeedbackModel struct {
	ID         uint32
	UserID     uint32
	Kind       int8 // 反馈类型：0产品建议，1功能异常，2违规举报，3交易投诉
	Content    string
	Pictures   string // 反馈图片，分号分割
	HasRead    bool
	CreateTime time.Time
	ReadTime   time.Time
}

func (m *FeedbackModel) TableName() string {
	return "feedback"
}

func (m *FeedbackModel) Create() error {
	m.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(m).Error
}

func (m *FeedbackModel) Save() error {
	return DB.Self.Save(m).Error
}

func GetFeedbacks(limit, offset int, kind, readFilter int8) ([]*FeedbackModel, error) {
	var list []*FeedbackModel
	query := DB.Self
	if kind != -1 {
		query = query.Where("kind = ?", kind)
	}
	if readFilter != -1 {
		query = query.Where("has_read = ?", readFilter)
	}

	d := query.Limit(limit).Offset(offset).Find(&list)
	if d.RecordNotFound() {
		return list, nil
	}
	return list, d.Error
}

func FeedbackBatchRead(ids []uint32) error {
	now := util.GetCurrentTime()
	err := DB.Self.Model(FeedbackModel{}).Where("id in (?)", ids).
		Updates(map[string]interface{}{"has_read": 1, "read_time": now}).
		Error
	return err
}
