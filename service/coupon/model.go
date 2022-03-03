package coupon

type BasicCoupon struct {
	ID        uint32 `json:"id"`
	Cid       uint32 `json:"cid"`        // 所属一级类目
	Cid2      uint32 `json:"cid_2"`      // 所属二级类目
	Discount  int    `json:"discount"`   // 折扣金额
	Threshold int    `json:"threshold"`  // 满减门槛
	Kind      int8   `json:"kind"`       // 优惠券种类，默认为0，暂时就一种
	Title     string `json:"title"`      // 标题
	BeginTime string `json:"begin_time"` // 生效开始时间
	EndTime   string `json:"end_time"`   // 生效结束时间
}

type PublicItem struct {
	BasicCoupon
	Remain        int    `json:"remain"`          // 剩余数
	GrabBeginTime string `json:"grab_begin_time"` // 领取开始时间
	GrabEndTime   string `json:"grab_end_time"`   // 领取截止时间
	HasGrabbed    bool   `json:"has_grabbed"`
}

type PrivateItem struct {
	BasicCoupon
	Access int8 `json:"access"`
}

// 优惠券设置item
type CouponConfigItem struct {
	BasicCoupon
	IsPublic      bool   `json:"is_public"`       // 是否公共可领取
	Code          string `json:"code"`            // 兑换码
	Remain        int    `json:"remain"`          // 剩余数
	GrabBeginTime string `json:"grab_begin_time"` // 领取开始时间
	GrabEndTime   string `json:"grab_end_time"`   // 领取截止时间
	CodeBeginTime string `json:"code_begin_time"` // 兑换开始时间
	CodeEndTime   string `json:"code_end_time"`   // 兑换截止时间
}

// 管理端查询列表item
type AdminItem struct {
	CouponConfigItem
	GrabNum    int64  `json:"grab_num"`    // 领取人数
	CreateTime string `json:"create_time"` // 记录/优惠券创建时间
}
