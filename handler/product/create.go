package product

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func CreateOrUpdate(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
