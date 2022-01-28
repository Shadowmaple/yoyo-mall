package collection

import (
	"log"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/collection"

	"github.com/gin-gonic/gin"
)

type ListReq struct {
	Limit int
	Page  int
}

func List(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover error:", err)
		}
	}()

	req := &ListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	list, err := collection.List(userID, req.Limit, req.Page)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, map[string]interface{}{
		"total": len(list),
		"list":  list,
	})
}
