package cart

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/cart"

	"github.com/gin-gonic/gin"
)

type AddReq struct {
	List []cart.BasicItem
}

// 添加商品
func Add(c *gin.Context) {
	req := &AddReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)
	if err := cart.BatchAdd(userID, req.List); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
