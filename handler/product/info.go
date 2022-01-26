package product

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
