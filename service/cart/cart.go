package cart

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

type BasicItem struct {
	ID  uint32 `json:"id"`
	Num int    `json:"num"`
}

// 批量插入
func BatchAdd(userID uint32, list []BasicItem) error {
	records := make([]*model.CartModel, 0, len(list))

	for _, item := range list {
		records = append(records, &model.CartModel{
			UserID:     userID,
			ProductID:  item.ID,
			Num:        item.Num,
			CreateTime: util.GetCurrentTime(),
		})
	}

	return model.CartBatchInsert(records)
}

// 批量更新
func BatchUpdate(userID uint32, list []BasicItem) error {
	for _, item := range list {
		if err := model.UpdateCartNum(item.ID, item.Num); err != nil {
			return err
		}
	}
	return nil
}

// 批量删除
// 有个漏洞，没有删除和修改的记录权限进行校验，即该用户是否能删除或修改
func BatchDelete(userID uint32, list []uint32) error {
	return model.CartBatchDelete(list)
}
