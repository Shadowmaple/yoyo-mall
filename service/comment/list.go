package comment

import "yoyo-mall/model"

func List(userID, evaluationID uint32, limit, page int) (list []*CommentItem, err error) {
	list = make([]*CommentItem, 0)

	records, err := model.GetComments(evaluationID, limit, limit*page)
	if err != nil {
		return
	}

	for _, record := range records {

		userNickname, userAvatar := "", ""
		if !record.IsAnoymous {
			user, err1 := model.GetUserByID(record.UserID)
			if err1 != nil {
				err = err1
				return
			}
			userNickname = user.Nickname
			userAvatar = user.Avatar
		}

		hasLiked, canHandle := false, false
		if userID > 0 {
			hasLiked = model.HasLiked(userID, record.ID, 1)
			if record.UserID == userID {
				canHandle = true
			}
		}

		likeNum, err1 := model.GetLikeNum(record.ID, 1)
		if err1 != nil {
			err = err1
			return
		}

		list = append(list, &CommentItem{
			BasicItem: BasicItem{
				ID:         record.ID,
				Content:    record.Content,
				IsAnoymous: record.IsAnoymous,
			},
			UserNickname: userNickname,
			UserAvatar:   userAvatar,
			LikeNum:      likeNum,
			HasLiked:     hasLiked,
			CanHandle:    canHandle,
		})
	}

	return
}
