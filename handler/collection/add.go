package collection

import (
	"log"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/collection"

	"github.com/gin-gonic/gin"
)

type AddReq struct {
	List []uint32 `json:"list"`
}

func Add(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover error:", err)
		}
	}()

	req := &AddReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	if err := collection.BatchAdd(userID, req.List); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
