package order

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
