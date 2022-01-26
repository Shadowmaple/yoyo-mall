package upload

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

// 图片上传
func Image(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
