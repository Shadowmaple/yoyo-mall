package model

import (
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

/*

订单状态：
0->待付款，1->待发货，2->待收货，3->待评价，4->交易完成，5->交易取消，6->退货中，7->交易关闭

状态变更图：
1. 正常交易：0待付款->1待发货->2待收货->3待评价->4交易完成
2. 未付款取消订单：0待付款->5交易取消
3. 退货：0待付款->……->6退货中->7交易关闭

*/

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

func (m *OrderModel) TableName() string {
	return "order"
}

func (m *OrderModel) Create() error {
	m.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(m).Error
}

func (m *OrderModel) Save() error {
	return DB.Self.Save(m).Error
}

func GetOrderByID(id uint32) (*OrderModel, error) {
	model := &OrderModel{}
	d := DB.Self.First(model, "id = ?", id)
	if d.RecordNotFound() {
		return model, errno.ErrRecordNotFound
	}
	return model, d.Error
}

func OrderList(userID uint32, limit, offset int, status int8) ([]*OrderModel, error) {
	list := make([]*OrderModel, 0)

	query := DB.Self.Where("user_id = ?", userID)

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	d := query.Limit(limit).Offset(offset).Order("id desc").Find(&list)

	if d.RecordNotFound() {
		return list, nil
	}

	return list, d.Error
}

func OrderSearch(sql string) ([]*OrderModel, error) {
	list := make([]*OrderModel, 0)
	err := DB.Self.Raw(sql).Scan(&list).Error
	return list, err
}
