package cart

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/cart"

	"github.com/gin-gonic/gin"
)

type UpdateReq struct {
	List []cart.BasicItem `json:"list"`
}

// 修改商品数量等信息
func Update(c *gin.Context) {
	req := &UpdateReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)
	if err := cart.BatchUpdate(userID, req.List); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
