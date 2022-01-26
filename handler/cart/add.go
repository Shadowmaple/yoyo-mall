package cart

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 添加商品
func Add(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
