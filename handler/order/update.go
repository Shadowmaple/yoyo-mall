package order

import (
	"fmt"
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/log"
	"yoyo-mall/service/order"

	"github.com/gin-gonic/gin"
)

type UpdateReq struct {
	Status int8 `form:"status"`
}

// 修改订单状态，只能为1/2/3/5/6，即完成付款、已发货、已签收、取消订单、退货
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Get param 'id' error: " + err.Error())
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	req := &UpdateReq{}
	if err := c.BindJSON(req); err != nil {
		log.Error("BindJson error: " + err.Error())
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	role := c.MustGet("role").(uint8)

	expectedStatus := req.Status
	switch expectedStatus {
	case 1, 3, 5, 6:
		if role != 0 {
			log.Error(fmt.Sprintf("Permission denied: role=%d, req=%+v", role, req))
			handler.SendBadRequest(c, errno.ErrOrderExpectedStatus, nil, "permission denied")
			return
		}
	case 2:
		if role != 1 {
			log.Error(fmt.Sprintf("Permission denied: role=%d, req=%+v", role, req))
			handler.SendBadRequest(c, errno.ErrOrderExpectedStatus, nil, "permission denied")
			return
		}
	}

	if err := order.UpdateStatus(uint32(id), expectedStatus); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
