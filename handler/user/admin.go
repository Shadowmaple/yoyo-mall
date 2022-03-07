package user

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/token"
	"yoyo-mall/service/user"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AdminLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AdminLogin(c *gin.Context) {
	req := AdminLoginReq{}
	if err := c.BindJSON(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if req.Username == "" || req.Password == "" {
		handler.SendBadRequest(c, errno.ErrBind, nil, "用户名和密码不能为空")
		return
	}

	id, err := user.AdminLogin(req.Username, req.Password)
	if err != nil {
		// 用户名错误，用户不存在
		if err == errno.ErrRecordNotFound {
			handler.SendResponse(c, errno.ErrUserNotFound, LoginResp{})
			return
		}
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	// 密码错误登录失败
	if id == 0 {
		handler.SendResponse(c, errno.ErrPwdWrong, LoginResp{})
		return
	}

	t, err := token.Sign(c, token.Context{ID: id, Role: 1}, viper.GetString("jwt_secret"))
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, map[string]interface{}{
		"token": t,
	})
}
