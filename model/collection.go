package model

import (
	"time"
	"yoyo-mall/util"
)

type ColletionModel struct {
	ID         uint32
	UserID     uint32
	ProductID  uint32
	CreateTime *time.Time
	IsDeleted  bool
	DeleteTime *time.Time
}

const CollectionTableName = "cart"

func (c *ColletionModel) TableName() string {
	return CollectionTableName
}

func (c *ColletionModel) Create() error {
	c.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(c).Error
}

func HasStar(userID, productID uint32) bool {
	m := &CartModel{}
	d := DB.Self.Table(CollectionTableName).
		Where("is_deleted = 0").
		Where("user_id = ? and product_id = ?", userID, productID).
		First(m)

	if d.RecordNotFound() {
		return false
	}
	return true
}

func UnStar(userID, productID uint32) error {
	deleteTime := util.GetStandardTime(util.GetCurrentTime())
	err := DB.Self.
		Where("is_deleted = 0").
		Where("user_id = ? and product_id = ?", userID, productID).
		Update(map[string]interface{}{"is_deleted": 1, "delete_time": deleteTime}).
		Error

	return err
}
