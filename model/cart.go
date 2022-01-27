package model

import "time"

// 购物车
type CartModel struct {
	ID         uint32
	UserID     uint32
	ProductID  uint32
	Num        int // 商品数量
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}

const CartTableName = "cart"

func (c *CartModel) TableName() string {
	return CartTableName
}

func GetProductNumInCart(userID, productID uint32) int {
	m := &CartModel{}
	d := DB.Self.Table(CartTableName).
		Where("is_deleted = 0").
		Where("user_id = ? and product_id = ?", userID, productID).
		First(m)

	if d.RecordNotFound() {
		return 0
	}
	return m.Num
}

func HasInCart(userID, productID uint32) bool {
	num := GetProductNumInCart(userID, productID)
	return num > 0
}
