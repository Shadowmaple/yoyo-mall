package handler

import (
	"net/http"
	"yoyo-mall/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	// log.Info(message, lager.Data{"X-Request-Id": util.GetReqID(c), "cause": cause})
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	// log.Info(message, lager.Data{"X-Request-Id": util.GetReqID(c), "cause": cause})
	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendUnauthorized(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	// log.Info(message, lager.Data{"X-Request-Id": util.GetReqID(c), "cause": cause})
	c.JSON(http.StatusUnauthorized, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}
