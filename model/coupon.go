package model

import (
	"errors"
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type CouponModel struct {
	ID            uint32
	Cid           uint32    // 所属一级类目
	Cid2          uint32    // 所属二级类目
	Discount      int       // 折扣金额
	Threshold     int       // 满减门槛
	Kind          int8      // 优惠券种类，默认为0，暂时就一种
	IsPublic      bool      // 是否公共可领取
	Code          string    // 兑换码
	Title         string    // 标题
	Remain        int       // 剩余数
	BeginTime     time.Time // 生效开始时间
	EndTime       time.Time // 生效结束时间
	GrabBeginTime time.Time // 领取开始时间
	GrabEndTime   time.Time // 领取截止时间
	CodeBeginTime time.Time // 兑换开始时间
	CodeEndTime   time.Time // 兑换截止时间
	CreateTime    time.Time
	IsDeleted     bool
	DeleteTime    *time.Time
}

func (m *CouponModel) TableName() string {
	return "coupon"
}

func (m *CouponModel) Create() error {
	m.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(m).Error
}

func (m *CouponModel) Save() error {
	return DB.Self.Save(m).Error
}

type CouponFilter struct {
	Cid    uint32
	Cid2   uint32
	Public bool
}

func GetCoupons(limit, offset int, filter *CouponFilter) ([]*CouponModel, error) {
	list := make([]*CouponModel, 0)
	now := util.GetCurrentTime()

	// 未删除的且未失效的
	query := DB.Self.Where("is_deleted = 0 and end_time > ?", now)

	if filter.Public {
		query = query.Where("is_public = 1")
	}

	if filter.Cid > 0 {
		query = query.Where("cid = ?", filter.Cid)
	}
	if filter.Cid2 > 0 {
		query = query.Where("cid2 = ?", filter.Cid2)
	}

	err := query.Limit(limit).Offset(offset).Find(&list).Error

	return list, err
}

func GetCouponByID(id uint32) (*CouponModel, error) {
	var m CouponModel
	err := DB.Self.Where("is_deleted = 0").Where("id = ?", id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errno.ErrRecordNotFound
	}
	return &m, err
}

func GetCouponByCode(code string) (CouponModel, error) {
	var m CouponModel
	err := DB.Self.Where("is_deleted = 0").Where("code = ?", code).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, errno.ErrRecordNotFound
	}
	return m, err
}

func DeleteCoupon(id uint32) error {
	now := util.GetCurrentTime()
	err := DB.Self.Model(CouponModel{}).Where("id = ?", id).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": now}).
		Error
	return err
}
