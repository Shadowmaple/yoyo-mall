package cart

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

type BasicItem struct {
	ID  uint32 `json:"id"` // 插入：product_id，修改：主键id
	Num int    `json:"num"`
}

// 批量插入
func BatchAdd(userID uint32, list []BasicItem) error {
	// 获取已有的记录，防止重复插入
	models, err := model.GetCarts(userID)
	if err != nil {
		return err
	}
	existProducts := make(map[uint32]bool, len(models))
	for _, model := range models {
		existProducts[model.ProductID] = true
	}

	records := make([]model.CartModel, 0, len(list))

	for _, item := range list {
		if _, ok := existProducts[item.ID]; ok {
			continue
		}
		if item.Num <= 0 {
			item.Num = 1
		}
		records = append(records, model.CartModel{
			UserID:     userID,
			ProductID:  item.ID,
			Num:        item.Num,
			CreateTime: util.GetCurrentTime(),
		})
	}
	if len(records) == 0 {
		return nil
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
