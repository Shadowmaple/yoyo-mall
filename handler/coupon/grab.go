package coupon

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 获取优惠券
func Grab(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
