package evaluation

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

func Create(userID uint32, req *BasicItem) (err error) {
	records := make([]*model.EvaluationModel, 0)

	productRecords, err := model.OrderProductModel{}.GetByOrderID(req.OrderID)
	if err != nil {
		return
	}

	now := util.GetCurrentTime()

	for _, item := range productRecords {
		productID := item.ProductID

		records = append(records, &model.EvaluationModel{
			UserID:     userID,
			OrderID:    req.OrderID,
			ProductID:  productID,
			Content:    req.Content,
			Score:      req.Score,
			Level:      req.Level,
			IsAnoymous: req.IsAnoymous,
			Pictures:   util.MergeMultiImage(req.Pictures),
			CreateTime: now,
		})
	}

	// 事务
	tx := model.DB.Self.Begin()

	// 创建评价
	if err = tx.Create(records).Error; err != nil {
		tx.Rollback()
		return
	}

	// 修改订单状态为已评价/交易完成
	if err = tx.Model(&model.OrderModel{}).Update("status", 4).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

func New(userID uint32, req *BasicItem) (err error) {

	records := make([]*model.EvaluationModel, 0)

	productRecords, err := model.OrderProductModel{}.GetByOrderID(req.OrderID)
	if err != nil {
		return
	}

	now := util.GetCurrentTime()

	for _, item := range productRecords {
		productID := item.ProductID

		records = append(records, &model.EvaluationModel{
			UserID:     userID,
			OrderID:    req.OrderID,
			ProductID:  productID,
			Content:    req.Content,
			Score:      req.Score,
			Level:      req.Level,
			IsAnoymous: req.IsAnoymous,
			Pictures:   util.MergeMultiImage(req.Pictures),
			CreateTime: now,
		})
	}

	if err = model.BatchCreateEvaluations(records); err != nil {
		return
	}

	return
}

// todo：多条同内容评价的bug
func Update(userID uint32, req *BasicItem) (err error) {
	record, err := model.GetEvaluationByID(req.ID)
	if err != nil {
		return
	}
	record.Content = req.Content
	record.Score = req.Score
	record.Level = req.Level
	record.IsAnoymous = req.IsAnoymous
	record.Pictures = util.MergeMultiImage(req.Pictures)

	if err = record.Save(); err != nil {
		return
	}

	return
}
