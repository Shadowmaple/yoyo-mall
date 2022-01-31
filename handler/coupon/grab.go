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

	userID := c.MustGet("id").(uint32)

	// 是領取还是兑换
	isGrab := false
	if len(req.Code) == 0 {
		isGrab = true
	}

	data, err := coupon.GrabCoupon(userID, req.ID, req.Code, isGrab)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, data)
}
