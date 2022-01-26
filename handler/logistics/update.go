package logistics

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 修改物流状态
func Update(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
