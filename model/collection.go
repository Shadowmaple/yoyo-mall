package model

import (
	"errors"
	"time"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type ColletionModel struct {
	ID         uint32    `gorm:"column:id"`
	UserID     uint32    `gorm:"column:user_id"`
	ProductID  uint32    `gorm:"column:product_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	IsDeleted  bool      `gorm:"is_deleted"`
	DeleteTime *time.Time
}

const CollectionTableName = "collection"

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

// 批量删除（根据记录id）
func CollectBatchDelete(list []uint32) error {
	now := util.GetCurrentTime()
	err := DB.Self.Model(ColletionModel{}).Where("id in (?)", list).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": now}).
		Error

	return err
}

// 根据商品id和用户id删除
func CollectDelByProductID(userID, productID uint32) error {
	now := util.GetCurrentTime()
	err := DB.Self.Model(ColletionModel{}).Where("user_id = ? and product_id = ?", userID, productID).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": now}).
		Error

	return err
}

func HasStar(userID, productID uint32) bool {
	m := &ColletionModel{}
	err := DB.Self.
		Where("is_deleted = 0").
		Where("user_id = ? and product_id = ?", userID, productID).
		First(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func GetCollection(userID uint32, limit, offset int) ([]*ColletionModel, error) {
	list := make([]*ColletionModel, 0)
	d := DB.Self.Where("is_deleted = 0").
		Where("user_id = ?", userID).
		Order("id desc").
		Limit(limit).Offset(offset).
		Find(&list)

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
