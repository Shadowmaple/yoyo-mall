package cart

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/cart"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	userID := c.MustGet("id").(uint32)
	list, err := cart.List(userID)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, map[string]interface{}{
		"total": len(list),
		"list":  list,
	})
}
