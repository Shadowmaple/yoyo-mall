package order

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/order"

	"github.com/gin-gonic/gin"
)

type ListReq struct {
	Kind  int `form:"kind"`
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

func List(c *gin.Context) {
	req := &ListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	userID := c.MustGet("id").(uint32)

	list, err := order.List(userID, req.Limit, req.Page, req.Kind)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, map[string]interface{}{
		"total": len(list),
		"list":  list,
	})
}

func AdminList(c *gin.Context) {
	req := &ListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	list, err := order.List(0, req.Limit, req.Page, req.Kind)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, map[string]interface{}{
		"total": len(list),
		"list":  list,
	})
}
