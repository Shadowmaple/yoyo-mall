package model

import (
	"errors"
	"time"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type UserCouponModel struct {
	ID         uint32
	UserID     uint32
	CouponID   uint32
	Status     int8      // 使用状态：0未使用，1已使用
	Access     int8      // 获取方式：0领取，1兑换码
	CreateTime time.Time // 获取时间
}

func (m *UserCouponModel) TableName() string {
	return "user_coupon"
}

func (m *UserCouponModel) Create() error {
	m.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(m).Error
}

func (m *UserCouponModel) Save() error {
	return DB.Self.Save(m).Error
}

func UpdateUserCouponStatus(userID, couponID uint32, status int8) error {
	err := DB.Self.Model(UserCouponModel{}).
		Where("user_id = ? and coupon_id = ?", userID, couponID).
		Update("status", status).
		Error

	return err
}

func HasGrabCoupon(userID, couponID uint32) bool {
	var m UserCouponModel
	d := DB.Self.Where("user_id = ? and coupon_id = ?", userID, couponID).First(&m)
	if errors.Is(d.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func CountCouponGrabNum(couponID uint32) (count int64) {
	DB.Self.Model(&UserCouponModel{}).
		Where("coupon_id = ?", couponID).
		Count(&count)

	return
}

type UserCouponItem struct {
	CouponModel
	Status int8
	Access int8
}

func GetUserCoupon(userID uint32, status int8) ([]*UserCouponItem, error) {
	list := make([]*UserCouponItem, 0)
	var err error
	now := util.GetStandardTime(util.GetCurrentTime())

	// 已过期的
	if status == 2 {
		err = DB.Self.Select("coupon.*, user_coupon.status, user_coupon.access").
			Table("user_coupon").
			Joins("left join coupon on coupon.id = user_coupon.coupon_id").
			Where("user_coupon.user_id = ?", userID).
			Where("coupon.end_time < ?", now).
			Where("coupon.is_deleted = 0").
			Find(&list).
			Error
	} else {
		err = DB.Self.Select("coupon.*, user_coupon.status, user_coupon.access").
			Table("user_coupon").
			Joins("left join coupon on coupon.id = user_coupon.coupon_id").
			Where("user_coupon.user_id = ?", userID).
			Where("user_coupon.status = ?", status).
			Where("coupon.end_time > ?", now).
			Where("coupon.is_deleted = 0").
			Find(&list).
			Error
	}

	return list, err
}
