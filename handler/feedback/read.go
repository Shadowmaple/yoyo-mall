package feedback

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/feedback"

	"github.com/gin-gonic/gin"
)

type ReadReq struct {
	Data []uint32
}

func Read(c *gin.Context) {
	req := &ReadReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if err := feedback.Read(req.Data); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
