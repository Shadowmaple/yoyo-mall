package upload

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util/upload"

	"github.com/gin-gonic/gin"
)

type ImageResp struct {
	URL string `json:"url"`
}

// 图片上传
func Image(c *gin.Context) {
	image, header, err := c.Request.FormFile("image")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetFile, nil, err.Error())
		return
	}
	dataLen := header.Size
	userID := c.MustGet("id").(uint32)

	url, err := upload.UploadImage(header.Filename, userID, image, dataLen)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, ImageResp{URL: url})
}
