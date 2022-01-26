package product

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 获取榜单
func GetRank(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
