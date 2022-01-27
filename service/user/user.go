package user

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"
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
