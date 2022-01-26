package model

import "time"

type OrderModel struct {
	ID          uint32
	UserID      uint32
	Status      int8    // 状态：0->待付款，1->待发货，2->待收货，3->待评价，4->交易完成，5->交易取消，6->退货中，7->交易关闭
	Payment     float32 // 实付金额
	Freight     float32 // 运费
	TotalFee    float32 // 总金额
	Coupon      float32 // 优惠金额
	ReceiveName string  // 收货人
	ReceiveTel  string
	ReceiveAddr string
	Refund      string // 退货退款内容，暂时占位
	OrderCode   string // 订单编号
	CreateTime  *time.Time
	PayTime     *time.Time // 付款时间
	DeliverTime *time.Time // 发货时间
	ConfirmTime *time.Time // 签收时间
}

type OrderProductModel struct {
	ID         uint32
	OrderID    uint32
	ProductID  uint32
	Num        int
	Price      float32 // 单价-原价
	CurPrice   float32 // 单价-优惠价
	TotalFee   float32 // 总金额
	Image      string  // 封面图片
	CreateTime *time.Time
}
