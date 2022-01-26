package cart

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 修改商品数量等信息
func Update(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
