package logistics

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/logistics"

	"github.com/gin-gonic/gin"
)

type UpdateReq struct {
	OrderID uint32 `json:"order_id"`
	State   int8   `json:"state"`
}

// 修改物流状态
func Update(c *gin.Context) {
	req := &UpdateReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if err := logistics.NewLogistics(req.OrderID, req.State); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
