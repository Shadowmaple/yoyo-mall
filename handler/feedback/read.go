package feedback

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
