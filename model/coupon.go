package model

import (
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

type CouponModel struct {
	ID            uint32
	Cid           uint32     // 所属一级类目
	Cid2          uint32     // 所属二级类目
	Discount      int        // 折扣金额
	Threshold     int        // 满减门槛
	Kind          int8       // 优惠券种类，默认为0，暂时就一种
	IsPublic      bool       // 是否公共可领取
	Code          string     // 兑换码
	Title         string     // 标题
	Remain        int        // 剩余数
	BeginTime     *time.Time // 生效开始时间
	EndTime       *time.Time // 生效结束时间
	GrabBeginTime *time.Time // 领取开始时间
	GrabEndTime   *time.Time // 领取截止时间
	CodeBeginTime *time.Time // 兑换开始时间
	CodeEndTime   *time.Time // 兑换截止时间
	CreateTime    *time.Time
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

func GetCoupons(limit, offset int, cid, cid2 uint32) ([]*CouponModel, error) {
	list := make([]*CouponModel, 0)
	now := util.GetCurrentTime()

	// 未删除的且未失效的
	query := DB.Self.Where("is_deleted = 0 and end_time > ?", now)

	if cid > 0 {
		query = query.Where("cid = ?", cid)
	}
	if cid2 > 0 {
		query = query.Where("cid2 = ?", cid2)
	}

	d := query.Limit(limit).Offset(offset).Find(&list)
	if d.RecordNotFound() {
		return list, nil
	}

	return list, d.Error
}

func GetCouponByID(id uint32) (*CouponModel, error) {
	var m *CouponModel
	d := DB.Self.Where("is_deleted = 0").Where("id = ?", id).First(m)
	if d.RecordNotFound() {
		return nil, errno.ErrRecordNotFound
	}
	return m, d.Error
}
