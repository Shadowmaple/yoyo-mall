package user

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/auth"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/log"
)

func Login(openID string) (id uint32, isNew bool, err error) {
	u, err := model.GetUserByWechat(openID)
	if err == errno.ErrRecordNotFound {
		isNew = true
		u.WechatUniqueID = openID
		u.Role = 0
		if err = u.Create(); err != nil {
			return
		}
		id = u.ID
		return
	} else if err != nil {
		return
	}

	id = u.ID
	isNew = false
	return
}

func AdminLogin(username, password string) (id uint32, err error) {
	u, err := model.GetUserByUsername(username)
	if err != nil {
		return
	}

	err = auth.Compare(u.Password, password)
	if err != nil {
		log.Info("auth.Compare error; login failed: " + err.Error())
		err = nil
		return
	}

	id = u.ID

	return
}
