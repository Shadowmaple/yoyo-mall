package coupon

import (
	"errors"
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

// 使用优惠券
func UseCoupon(userID, couponID uint32) (err error) {
	err = model.UpdateUserCouponStatus(userID, couponID, 1)
	return
}

// 兑换码兑换
func GrabCouponByCode(userID uint32, code string) (item PrivateItem, err error) {
	// 查找优惠券
	coupon, err := model.GetCouponByCode(code)
	if err != nil {
		if err == errno.ErrRecordNotFound {
			err = errno.ErrCouponCodeWrong
			return
		}
		return
	}

	// 是否已领取
	if model.HasGrabCoupon(userID, coupon.ID) {
		err = errno.ErrCouponGrabbed
		return
	}

	// 现在是否处于可兑换时间
	now := util.GetCurrentTime()
	if now.Before(coupon.CodeBeginTime) || now.After(coupon.CodeEndTime) {
		err = errno.ErrCouponCanNotGrab
		return
	}

	record := &model.UserCouponModel{
		UserID:   userID,
		CouponID: coupon.ID,
		Status:   0,
		Access:   1,
	}
	if err = record.Create(); err != nil {
		return
	}

	item = PrivateItem{
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

// 领取/兑换优惠券
func GrabCoupon(userID, couponID uint32) (item PrivateItem, err error) {
	if model.HasGrabCoupon(userID, couponID) {
		err = errors.New("has grabbed")
		return
	}
	coupon, err := model.GetCouponByID(couponID)
	if err != nil {
		if err == errno.ErrRecordNotFound {
			err = errno.ErrCouponNotExist
			return
		}
		return
	}
	if !coupon.IsPublic {
		err = errno.ErrCouponNotPublic
		return
	}

	// 现在是否处于可领取时间
	now := util.GetCurrentTime()
	if now.Before(coupon.GrabBeginTime) || now.After(coupon.GrabEndTime) {
		err = errno.ErrCouponCanNotGrab
		return
	}

	record := &model.UserCouponModel{
		UserID:   userID,
		CouponID: couponID,
		Status:   0,
		Access:   0,
	}
	if err = record.Create(); err != nil {
		return
	}

	item = PrivateItem{
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
