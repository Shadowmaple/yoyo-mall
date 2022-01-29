package feedback

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

func New(userID uint32, req *BasicItem) (err error) {
	record := &model.FeedbackModel{
		UserID:   userID,
		Kind:     req.Kind,
		Content:  req.Content,
		Pictures: util.MergeMultiImage(req.Pictures),
	}

	if err = record.Create(); err != nil {
		return
	}

	return
}

func Read(data []uint32) (err error) {
	return model.FeedbackBatchRead(data)
}
