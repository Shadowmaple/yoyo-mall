package coupon

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func PrivateList(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}

func PublicList(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
