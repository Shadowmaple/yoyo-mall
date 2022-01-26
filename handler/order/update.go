package order

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 修改订单状态
func Update(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
