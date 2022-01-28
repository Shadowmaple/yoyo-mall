package order

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/order"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	data, err := order.GetInfo(uint32(id))
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, data)
}
