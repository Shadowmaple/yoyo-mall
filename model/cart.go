package model

import (
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

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

func (c *CartModel) Create() error {
	c.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(c).Error
}

func CartBatchInsert(records []*CartModel) error {
	return DB.Self.Create(&records).Error
}

func UpdateCartNum(id uint32, num int) error {
	err := DB.Self.
		Where("is_deleted = 0").
		Where("id = ?", id).
		Update("num", num).Error

	return err
}

// func DeleteCartProduct(userID, productID uint32) error {
// 	deleteTime := util.GetStandardTime(util.GetCurrentTime())
// 	err := DB.Self.
// 		Where("is_deleted = 0").
// 		Where("user_id = ? and product_id = ?", userID, productID).
// 		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": deleteTime}).
// 		Error

// 	return err
// }

func CartBatchDelete(list []uint32) error {
	now := util.GetCurrentTime()
	err := DB.Self.Model(CartModel{}).Where("id in ?", list).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": now}).
		Error

	return err
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

func GetCarts(userID uint32) ([]*CartModel, error) {
	var list []*CartModel
	d := DB.Self.Where("is_deleted = 0").Where("user_id = ?", userID).Find(&list)
	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}
	return list, d.Error
}
