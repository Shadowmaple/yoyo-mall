package product

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	if err := product.Delete(uint32(id)); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
