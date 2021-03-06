package collection

import "yoyo-mall/model"

// 批量添加（收藏）
func BatchAdd(userID uint32, productIDs []uint32) error {
	for _, productID := range productIDs {
		if model.HasStar(userID, productID) {
			continue
		}
		record := &model.ColletionModel{
			UserID:    userID,
			ProductID: productID,
		}
		if err := record.Create(); err != nil {
			return err
		}
	}

	return nil
}

// 批量删除(取消收藏)，list中id是记录id
func BatchDelete(userID uint32, list []uint32) error {
	if err := model.CollectBatchDelete(list); err != nil {
		return err
	}
	return nil
}

func DelByProductID(userID uint32, productID uint32) error {
	if err := model.CollectDelByProductID(userID, productID); err != nil {
		return err
	}
	return nil
}
