package collection

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
