package cart

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/cart"

	"github.com/gin-gonic/gin"
)

type DelReq struct {
	List []uint32 `json:"list"`
}

// 删除商品
func Delete(c *gin.Context) {
	req := &DelReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)
	if err := cart.BatchDelete(userID, req.List); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
