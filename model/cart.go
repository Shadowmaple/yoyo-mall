package model

import (
	"errors"
	"time"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

// 购物车
type CartModel struct {
	ID         uint32
	UserID     uint32    `gorm:"column:user_id"`
	ProductID  uint32    `gorm:"column:product_id"`
	Num        int       `gorm:"column:num"` // 商品数量
	CreateTime time.Time `gorm:"column:create_time"`
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

func CartBatchInsert(records []CartModel) error {
	return DB.Self.Create(&records).Error
}

func UpdateCartNum(id uint32, num int) error {
	err := DB.Self.
		Model(CartModel{}).
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
	err := DB.Self.Model(CartModel{}).Where("id in (?)", list).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": now}).
		Error

	return err
}

func GetProductNumInCart(userID, productID uint32) int {
	m := &CartModel{}
	err := DB.Self.
		Where("is_deleted = 0").
		Where("user_id = ? and product_id = ?", userID, productID).
		First(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
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
	err := DB.Self.Where("is_deleted = 0").
		Where("user_id = ?", userID).Order("id desc").
		Find(&list).Error
	return list, err
}
