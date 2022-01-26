package order

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
