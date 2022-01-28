package order

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/order"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	userID := c.MustGet("id").(uint32)

	req := &order.NewOrderItem{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	orderID, err := order.New(userID, req)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, map[string]uint32{"id": orderID})
}
