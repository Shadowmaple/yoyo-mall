package model

import (
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

type AddressModel struct {
	ID         uint32
	UserID     uint32
	Name       string // 收货人姓名
	Tel        string
	Province   string // 省
	City       string // 市
	District   string // 区县
	Detail     string // 街道详情
	IsDefault  bool   // 是否默认地址
	CreateTime time.Time
	IsDeleted  bool
	DeleteTime time.Time
}

func (a *AddressModel) TableName() string {
	return "address"
}

func (a *AddressModel) Create() error {
	a.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(a).Error
}

func (a *AddressModel) Save() error {
	return DB.Self.Save(a).Error
}

func GetAddressByID(id uint32) (*AddressModel, error) {
	var model *AddressModel
	d := DB.Self.Where("is_deleted = 0").First(model, "id = ?", id)
	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}
	return model, d.Error
}

func AddressList(userID uint32) ([]*AddressModel, error) {
	list := make([]*AddressModel, 0)

	d := DB.Self.Where("is_deleted = 0").Where("user_id = ?", userID).Find(&list)
	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}

	return list, d.Error
}

// 修改非默认地址的is_default字段
func UpdateNotDefaultAddress(userID, id uint32) error {
	err := DB.Self.Model(AddressModel{}).
		Where("is_deleted = 0").
		Where("user_id = ? and id != ?", userID, id).
		Update("is_default", 0).
		Error

	return err
}

func DeleteAddress(userID, id uint32) error {
	deleteTime := util.GetStandardTime(util.GetCurrentTime())
	err := DB.Self.
		Where("is_deleted = 0").
		Where("user_id = ? and id = ?", userID, id).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": deleteTime}).
		Error

	return err
}
