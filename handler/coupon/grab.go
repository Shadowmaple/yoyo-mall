package coupon

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/coupon"

	"github.com/gin-gonic/gin"
)

type GrabReq struct {
	Code string `form:"code"`
	ID   uint32 `form:"id"`
}

// 获取优惠券
func Grab(c *gin.Context) {
	req := &GrabReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	if req.ID == 0 && req.Code == "" {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "The id and code are both empty.")
		return
	}

	var data coupon.PrivateItem
	var err error
	userID := c.MustGet("id").(uint32)

	// 是領取还是兑换
	// 领取
	if len(req.Code) == 0 {
		data, err = coupon.GrabCoupon(userID, req.ID)
	} else {
		data, err = coupon.GrabCouponByCode(userID, req.Code)
	}

	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, data)
}
