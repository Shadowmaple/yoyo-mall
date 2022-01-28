package search

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/order"

	"github.com/gin-gonic/gin"
)

type OrderSearchReq struct {
	Limit int
	Page  int
	Key   string
}

func OrderSearch(c *gin.Context) {
	req := &OrderSearchReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	list, err := order.Search(userID, req.Key)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}
