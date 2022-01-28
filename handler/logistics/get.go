package logistics

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/logistics"

	"github.com/gin-gonic/gin"
)

// 某订单物流信息
func Get(c *gin.Context) {
	orderIDStr := c.DefaultQuery("order_id", "")
	if len(orderIDStr) == 0 {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "no order_id")
		return
	}

	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "order_id is wrong")
		return
	}

	list, err := logistics.GetInfoByOrderID(uint32(orderID))
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}
