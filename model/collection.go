package model

import (
	"time"
	"yoyo-mall/pkg/errno"
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

// 批量插入
func CollectBatchInsert(records []*ColletionModel) error {
	return DB.Self.Create(&records).Error
}

// 批量删除
func CollectBatchDelete(list []uint32) error {
	now := util.GetCurrentTime()
	err := DB.Self.Model(ColletionModel{}).Where("id in ?", list).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": now}).
		Error

	return err
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

func GetCollection(userID uint32, limit, offset int) ([]*ColletionModel, error) {
	list := make([]*ColletionModel, 0)
	d := DB.Self.Where("is_deleted = 0").
		Where("user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&list)

	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}

	return list, d.Error
}

// func CollectDelete(userID, productID uint32) error {
// 	deleteTime := util.GetStandardTime(util.GetCurrentTime())
// 	err := DB.Self.
// 		Where("is_deleted = 0").
// 		Where("user_id = ? and product_id = ?", userID, productID).
// 		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": deleteTime}).
// 		Error

// 	return err
// }
