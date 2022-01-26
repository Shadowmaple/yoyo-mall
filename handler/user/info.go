package user

import (
	"yoyo-mall/handler"

	"github.com/gin-gonic/gin"
)

func UpdateInfo(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}

func GetInfo(c *gin.Context) {
	handler.SendResponse(c, nil, "ok")
}
