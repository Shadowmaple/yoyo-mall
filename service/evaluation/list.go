package evaluation

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

// 查找某商品的评价列表
// todo：一次sql
func List(userID, productID uint32, limit, page int) (list []*EvaluationItem, err error) {
	list = make([]*EvaluationItem, 0)

	records, err := model.GetEvaluationList(0, 0, productID, limit, limit*page)
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
			hasLiked = model.HasLiked(userID, record.ID, 0)
			if record.UserID == userID {
				canHandle = true
			}
		}

		likeNum, err1 := model.GetLikeNum(record.ID, 0)
		if err1 != nil {
			err = err1
			return
		}

		replyNum := model.CountComment(record.ID)

		list = append(list, &EvaluationItem{
			BasicItem: BasicItem{
				ID:         record.ID,
				OrderID:    record.OrderID,
				Content:    record.Content,
				Score:      record.Score,
				Level:      record.Level,
				IsAnoymous: record.IsAnoymous,
				Pictures:   util.ParseMultiImage(record.Pictures),
			},
			Time:         util.GetStandardTime(record.CreateTime),
			UserNickname: userNickname,
			UserAvatar:   userAvatar,
			LikeNum:      likeNum,
			ReplyNum:     replyNum,
			HasLiked:     hasLiked,
			CanHandle:    canHandle,
		})
	}

	return
}
