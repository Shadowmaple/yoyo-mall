package feedback

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/feedback"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	readStr := c.DefaultQuery("read", "-1")
	read, err := strconv.Atoi(readStr)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "wrong read")
		return
	}
	kindStr := c.DefaultQuery("kind", "-1")
	kind, err := strconv.Atoi(kindStr)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "wrong kind")
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "wrong limit")
		return
	}

	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "wrong page")
		return
	}

	list, err := feedback.List(limit, page, int8(kind), int8(read))
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}
