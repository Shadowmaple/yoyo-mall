package coupon

import (
	"errors"
	"yoyo-mall/model"
	"yoyo-mall/util"
)

// 使用优惠券
func UseCoupon(userID, couponID uint32) (err error) {
	err = model.UpdateUserCouponStatus(userID, couponID, 1)
	return
}

// 领取/兑换优惠券
func GrabCoupon(userID, couponID uint32, code string, isGrab bool) (item *PrivateItem, err error) {
	if model.HasGrabCoupon(userID, couponID) {
		err = errors.New("has grabbed")
		return
	}
	coupon, err := model.GetCouponByID(couponID)
	if err != nil {
		return
	}
	if isGrab && !coupon.IsPublic || !isGrab && coupon.Code != code {
		err = errors.New("grab failed")
		return
	}

	var access int8
	if !isGrab {
		access = 1
	}
	record := &model.UserCouponModel{
		UserID:   userID,
		CouponID: couponID,
		Status:   0,
		Access:   access,
	}
	if err = record.Create(); err != nil {
		return
	}

	item = &PrivateItem{
		BasicCoupon: BasicCoupon{
			ID:        coupon.ID,
			Cid:       coupon.Cid,
			Cid2:      coupon.Cid2,
			Discount:  coupon.Discount,
			Threshold: coupon.Threshold,
			Kind:      coupon.Kind,
			Title:     coupon.Title,
			BeginTime: util.GetStandardTime(coupon.BeginTime),
			EndTime:   util.GetStandardTime(coupon.EndTime),
		},
		Access: record.Access,
	}

	return
}

func AddCoupon(req *CouponConfigItem) (err error) {
	return
}

func UpdateCoupon(req *CouponConfigItem) (err error) {
	return
}
