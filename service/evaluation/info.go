package evaluation

import (
	"yoyo-mall/model"
)

// todo: 一次sql
func Info(id uint32, userID uint32) (res EvaluationItem, err error) {
	record, err := model.GetEvaluationByID(id)
	if err != nil {
		return
	}

	res, err = processRecord(record, userID)
	if err != nil {
		return
	}

	return
}
