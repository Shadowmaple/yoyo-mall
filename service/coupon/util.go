package coupon

import (
	"time"
	"yoyo-mall/util"
)

type TimeItem struct {
	BeginTime     time.Time
	EndTime       time.Time
	GrabBeginTime time.Time
	GrabEndTime   time.Time
	CodeBeginTime time.Time
	CodeEndTime   time.Time
}

func BatchParseTime(req *CouponConfigItem) (t TimeItem, err error) {
	t.BeginTime, err = util.ParseTime(req.BeginTime)
	if err != nil {
		return
	}
	t.EndTime, err = util.ParseTime(req.EndTime)
	if err != nil {
		return
	}
	t.GrabBeginTime, err = util.ParseTime(req.GrabBeginTime)
	if err != nil {
		return
	}
	t.GrabEndTime, err = util.ParseTime(req.GrabEndTime)
	if err != nil {
		return
	}
	t.CodeBeginTime, err = util.ParseTime(req.CodeBeginTime)
	if err != nil {
		return
	}
	t.CodeEndTime, err = util.ParseTime(req.CodeEndTime)
	if err != nil {
		return
	}
	return
}
