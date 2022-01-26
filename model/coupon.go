package model

import "time"

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

type UserCouponModel struct {
	ID         uint32
	UserID     uint32
	CouponID   uint32
	Status     int8       // 使用状态：0未使用，1已使用
	Access     int8       // 获取方式：0领取，1兑换码
	CreateTime *time.Time // 获取时间
}
