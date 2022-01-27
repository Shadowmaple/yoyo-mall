package model

import "time"

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
