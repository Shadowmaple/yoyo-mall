package user

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
