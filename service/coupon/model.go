package coupon

type BasicCoupon struct {
	ID        uint32
	Cid       uint32 // 所属一级类目
	Cid2      uint32 // 所属二级类目
	Discount  int    // 折扣金额
	Threshold int    // 满减门槛
	Kind      int8   // 优惠券种类，默认为0，暂时就一种
	Title     string // 标题
	BeginTime string // 生效开始时间
	EndTime   string // 生效结束时间
}

type PublicItem struct {
	BasicCoupon
	Remain        int    // 剩余数
	GrabBeginTime string // 领取开始时间
	GrabEndTime   string // 领取截止时间
	HasGrabbed    bool
}

type PrivateItem struct {
	BasicCoupon
	Access int8
}

type CouponConfigItem struct {
	BasicCoupon
	IsPublic      bool   // 是否公共可领取
	Code          string // 兑换码
	Remain        int    // 剩余数
	GrabBeginTime string // 领取开始时间
	GrabEndTime   string // 领取截止时间
	CodeBeginTime string // 兑换开始时间
	CodeEndTime   string // 兑换截止时间
}
