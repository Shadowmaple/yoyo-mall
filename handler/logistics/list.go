package logistics

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 管理端物流列表
func List(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
