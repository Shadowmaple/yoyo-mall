package product

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	limitStr := c.DefaultQuery("comment_limit", "2")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "comment_limit is error")
		return
	}

	var userID uint32
	idTemp, ok := c.Get("id")
	if ok {
		userID = idTemp.(uint32)
	}

	info, err := product.GetProfile(uint32(id), userID, limit)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, info)
}
