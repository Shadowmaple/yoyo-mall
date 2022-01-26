package comment

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func EvaluationList(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}

func EvaluationCreateOrUpdate(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
