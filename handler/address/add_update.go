package address

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func AddOrUpdate(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
