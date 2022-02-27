package coupon

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/log"
	"yoyo-mall/util"
)

func PrivateList(userID uint32, status int8) (list []*PrivateItem, err error) {
	list = make([]*PrivateItem, 0)

	records, err := model.GetUserCoupon(userID, status)
	if err != nil {
		return
	}

	for _, record := range records {
		list = append(list, &PrivateItem{
			BasicCoupon: BasicCoupon{
				ID:        record.ID,
				Cid:       record.Cid,
				Cid2:      record.Cid2,
				Discount:  record.Discount,
				Threshold: record.Threshold,
				Kind:      record.Kind,
				Title:     record.Title,
				BeginTime: util.GetStandardTime(record.BeginTime),
				EndTime:   util.GetStandardTime(record.EndTime),
			},
			Access: record.Access,
		})
	}

	return
}

// todo: 一次sql
func PublicList(userID uint32, page, limit int, cid, cid2 uint32) (list []*PublicItem, err error) {
	list = make([]*PublicItem, 0)

	records, err := model.GetCoupons(limit, limit*page, cid, cid2)
	if err != nil {
		log.Info("model.GetCoupons error:" + err.Error())
		return
	}

	for _, record := range records {
		hasGrabbed := model.HasGrabCoupon(userID, record.ID)

		list = append(list, &PublicItem{
			BasicCoupon: BasicCoupon{
				ID:        record.ID,
				Cid:       record.Cid,
				Cid2:      record.Cid2,
				Discount:  record.Discount,
				Threshold: record.Threshold,
				Kind:      record.Kind,
				Title:     record.Title,
				BeginTime: util.GetStandardTime(record.BeginTime),
				EndTime:   util.GetStandardTime(record.EndTime),
			},
			Remain:        record.Remain,
			GrabBeginTime: util.GetStandardTime(record.GrabBeginTime),
			GrabEndTime:   util.GetStandardTime(record.GrabEndTime),
			HasGrabbed:    hasGrabbed,
		})
	}

	return
}
