package cart

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 删除商品
func Delete(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
