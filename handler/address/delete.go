package address

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/address"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	if len(idStr) == 0 {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "no id")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "id error")
		return
	}

	userID := c.MustGet("id").(uint32)

	if err := address.Delete(userID, uint32(id)); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
