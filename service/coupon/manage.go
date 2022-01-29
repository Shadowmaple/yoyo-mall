package coupon

import "yoyo-mall/model"

func New(req *CouponConfigItem) (err error) {
	// 解析时间
	t, err := BatchParseTime(req)
	if err != nil {
		return
	}

	coupon := &model.CouponModel{
		Cid:           req.Cid,
		Cid2:          req.Cid2,
		Discount:      req.Discount,
		Threshold:     req.Threshold,
		Kind:          req.Kind,
		IsPublic:      req.IsPublic,
		Code:          req.Code,
		Title:         req.Title,
		Remain:        req.Remain,
		BeginTime:     t.BeginTime,
		EndTime:       t.EndTime,
		GrabBeginTime: t.GrabBeginTime,
		GrabEndTime:   t.GrabEndTime,
		CodeBeginTime: t.CodeBeginTime,
		CodeEndTime:   t.CodeEndTime,
	}

	if err = coupon.Create(); err != nil {
		return
	}

	return
}

func Update(req *CouponConfigItem) (err error) {
	coupon, err := model.GetCouponByID(req.ID)
	if err != nil {
		return
	}

	// 解析时间
	t, err := BatchParseTime(req)
	if err != nil {
		return
	}

	coupon = &model.CouponModel{
		ID:            coupon.ID,
		Cid:           req.Cid,
		Cid2:          req.Cid2,
		Discount:      req.Discount,
		Threshold:     req.Threshold,
		Kind:          req.Kind,
		IsPublic:      req.IsPublic,
		Code:          req.Code,
		Title:         req.Title,
		Remain:        req.Remain,
		BeginTime:     t.BeginTime,
		EndTime:       t.EndTime,
		GrabBeginTime: t.GrabBeginTime,
		GrabEndTime:   t.GrabEndTime,
		CodeBeginTime: t.CodeBeginTime,
		CodeEndTime:   t.CodeEndTime,
		IsDeleted:     coupon.IsDeleted,
		DeleteTime:    coupon.DeleteTime,
	}

	if err = coupon.Save(); err != nil {
		return
	}

	return
}

func Delete(id uint32) error {
	return model.DeleteCoupon(id)
}
