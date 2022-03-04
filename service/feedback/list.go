package feedback

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

func List(limit, page int, kind, read int8) (list []*FeedbackItem, err error) {
	list = make([]*FeedbackItem, 0)

	records, err := model.GetFeedbacks(limit, limit*page, kind, read)
	if err != nil {
		return
	}

	for _, record := range records {
		user, err1 := model.GetUserByID(record.UserID)
		if err1 != nil {
			err = err1
			return
		}

		list = append(list, &FeedbackItem{
			BasicItem: BasicItem{
				Kind:     record.Kind,
				Content:  record.Content,
				Pictures: util.ParseMultiImage(record.Pictures),
			},
			ID:           record.ID,
			Time:         util.GetStandardTime(record.CreateTime),
			HasRead:      record.HasRead,
			UserNickname: user.Nickname,
			UserAvatar:   user.Avatar,
		})
	}

	return
}
