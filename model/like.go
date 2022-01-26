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
