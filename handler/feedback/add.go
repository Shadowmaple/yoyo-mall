package feedback

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/feedback"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	req := &feedback.BasicItem{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	if err := feedback.New(userID, req); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
