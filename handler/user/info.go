package user

import (
	"yoyo-mall/handler"
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Gender   int8   `json:"gender"`
}

func UpdateInfo(c *gin.Context) {
	userID := c.MustGet("id").(uint32)

	req := &UserInfo{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	u, err := model.GetUserByID(userID)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	u.Avatar = req.Avatar
	u.Nickname = req.Nickname
	u.Gender = req.Gender
	if err = u.Save(); err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}

func GetInfo(c *gin.Context) {
	userID := c.MustGet("id").(uint32)

	u, err := model.GetUserByID(userID)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	resp := &UserInfo{
		Avatar:   u.Avatar,
		Nickname: u.Nickname,
		Gender:   u.Gender,
	}

	handler.SendResponse(c, nil, resp)
}
