package address

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/address"

	"github.com/gin-gonic/gin"
)

func AddOrUpdate(c *gin.Context) {
	req := &address.AddressInfo{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	var err error
	if req.ID == 0 {
		err = address.Add(userID, req)
	} else {
		err = address.Update(userID, req)
	}

	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
