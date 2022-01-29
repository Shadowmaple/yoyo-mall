package comment

import (
	"yoyo-mall/model"
)

func Publish(userID, evaluationID uint32, req *BasicItem) (err error) {
	record := &model.CommentModel{
		UserID:       userID,
		EvaluationID: evaluationID,
		Content:      req.Content,
		IsAnoymous:   req.IsAnoymous,
	}

	if err = record.Create(); err != nil {
		return
	}

	return
}

func Update(userID, evaluationID uint32, req *BasicItem) (err error) {
	record, err := model.GetCommentByID(req.ID)
	if err != nil {
		return
	}

	record.Content = req.Content
	record.IsAnoymous = req.IsAnoymous

	if err = record.Save(); err != nil {
		return
	}
	return
}
