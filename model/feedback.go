package model

import "time"

type Feedback struct {
	ID         uint32
	UserID     uint32
	Kind       int8 // 反馈类型：0产品建议，1功能异常，2违规举报，3交易投诉
	Content    string
	Pictures   string // 反馈图片，分号分割
	HasRead    bool
	CreateTime *time.Time
	ReadTime   *time.Time
}
