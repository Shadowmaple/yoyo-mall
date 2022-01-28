package model

import (
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

/*
物流状态：
0->待发货，1->已发货待收货，2->已完成，3->已取消，4->退货中，5->退货完成
*/
type LogisticsModel struct {
	ID         uint32
	OrderID    uint32
	Status     int8   // 状态：0->待发货，1->已发货待收货，2->已完成，3->已取消，4->退货中，5->退货完成
	Content    string // 物流信息
	CreateTime *time.Time
}

func (m *LogisticsModel) TableName() string {
	return "logistics"
}

func (m *LogisticsModel) Create() error {
	m.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(m).Error
}

func GetLogisticByID(id uint32) (*LogisticsModel, error) {
	model := &LogisticsModel{}
	d := DB.Self.First(model, "id = ?")
	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}
	return model, d.Error
}

func GetLogisticsByOrderID(orderID uint32) ([]*LogisticsModel, error) {
	list := make([]*LogisticsModel, 0)

	d := DB.Self.Where("order_id = ?", orderID).Find(&list)
	if d.RecordNotFound() {
		return list, errno.ErrRecordNotFound
	}
	return list, d.Error
}
