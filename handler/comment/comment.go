package comment

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func CommentList(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}

func CommentCreateOrUpdate(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
