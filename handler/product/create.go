package product

import (
	"fmt"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

func CreateOrUpdate(c *gin.Context) {
	req := &product.HandleItem{}
	if err := c.BindJSON(req); err != nil {
		fmt.Println("--", err.Error())
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	var err error
	if req.ID == 0 {
		err = product.New(req)
	} else {
		err = product.Update(req)
	}

	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
