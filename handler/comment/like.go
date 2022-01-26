package comment

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
