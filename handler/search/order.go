package search

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func OrderSearch(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
