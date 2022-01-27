package model

import "time"

type LikeModel struct {
	ID         uint32
	UserID     uint32
	CommentID  uint32 // 评价/评论id
	Kind       int8   // 0评价点赞，1评论点赞
	CreateTime *time.Time
	IsDeleted  bool
}

func (m *LikeModel) TableName() string {
	return "like"
}

func HasLiked(userID, commentID uint32, kind int8) bool {
	var count int8
	DB.Self.Where("is_deleted = 0").
		Where("user_id = ? and comment_id = ? and kind = ?", userID, commentID, kind).
		Count(&count)

	return count > 0
}
