package coupon

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/log"
	"yoyo-mall/service/coupon"

	"github.com/gin-gonic/gin"
)

func AddOrUpdate(c *gin.Context) {
	req := &coupon.CouponConfigItem{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 校验时间
	// ...

	var err error
	if req.ID == 0 {
		err = coupon.New(req)
	} else {
		err = coupon.Update(req)
	}
	if err != nil {
		log.Error("couponAddOrUpdate error:" + err.Error())
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}

func valid(req *coupon.CouponConfigItem) {

}
