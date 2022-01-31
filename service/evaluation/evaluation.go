package evaluation

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

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
