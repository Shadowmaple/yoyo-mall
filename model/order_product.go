package model

import "time"

type OrderProductModel struct {
	ID         uint32
	OrderID    uint32
	ProductID  uint32
	Num        int
	Price      float32 // 单价-原价
	CurPrice   float32 // 单价-优惠价
	TotalFee   float32 // 总金额
	Image      string  // 封面图片
	CreateTime time.Time
}

func (m *OrderProductModel) TableName() string {
	return "order_product"
}

type OrderProductInfo struct {
	OrderProductModel
	Title  string
	Author string
}

func GetProductByOrderID(orderID uint32) ([]*OrderProductInfo, error) {
	list := make([]*OrderProductInfo, 0)

	d := DB.Self.Table("order_product").Select("order_product.*, product.title, product.author").
		Joins("left join product on order_product.product_id = product.id").
		Where("order_id = ?", orderID).
		Find(&list)

	if d.RecordNotFound() {
		return list, nil
	}
	return list, d.Error
}

type OrderSearchItem struct {
	ID         uint32
	Status     int8
	TotalFee   float32
	ProductNum int
	ProductID  uint32
}
