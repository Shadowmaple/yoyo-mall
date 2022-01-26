package model

import "time"

type LogisticsModel struct {
	ID         uint32
	OrderID    uint32
	Status     int8   // 状态：0->待发货，1->已发货待收货，2->已完成，3->已取消，4->退货中，5->退货完成
	Content    string // 物流信息
	CreateTime *time.Time
}
