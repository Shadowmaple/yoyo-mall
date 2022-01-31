package model

import (
	"errors"
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type LikeModel struct {
	ID         uint32
	UserID     uint32
	CommentID  uint32 // 评价/评论id
	Kind       int8   // 0评价点赞，1评论点赞
	CreateTime time.Time
	IsDeleted  bool
}

func (m *LikeModel) TableName() string {
	return "like"
}

func HasLiked(userID, commentID uint32, kind int8) bool {
	var count int64
	DB.Self.Where("is_deleted = 0").
		Where("user_id = ? and comment_id = ? and kind = ?", userID, commentID, kind).
		Count(&count)

	return count > 0
}

func GetLikedRecord(userID, commentID uint32, kind, isDeleted int8) (*LikeModel, error) {
	var m *LikeModel
	d := DB.Self.Where("is_deleted = ?", isDeleted).
		Where("user_id = ? and comment_id = ? and kind = ?", userID, commentID, kind).
		First(m)

	if errors.Is(d.Error, gorm.ErrRecordNotFound) {
		return nil, errno.ErrRecordNotFound
	}

	return m, d.Error
}

func Unlike(userID, commentID uint32, kind int8) error {
	deleteTime := util.GetStandardTime(util.GetCurrentTime())
	err := DB.Self.
		Where("is_deleted = 0").
		Where("user_id = ? and comment_id = ? and kind = ?", userID, commentID, kind).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": deleteTime}).
		Error

	return err
}

// 点赞，目前是复用已删除的记录
func Like(userID, commentID uint32, kind int8) error {
	m, err := GetLikedRecord(userID, commentID, kind, 1)
	if err != nil {
		return err
	}
	// 存在删除的记录，修改
	if m != nil {
		m.IsDeleted = false
		m.CreateTime = util.GetCurrentTime()
		return DB.Self.Save(m).Error
	}

	// 无删除的记录，新建
	m = &LikeModel{
		UserID:     userID,
		CommentID:  commentID,
		Kind:       kind,
		CreateTime: util.GetCurrentTime(),
	}
	return DB.Self.Create(m).Error
}

func GetLikeNum(commentID uint32, kind int8) (int, error) {
	var count int64
	DB.Self.Where("is_deleted = 0").
		Where("comment_id = ? and kind = ?", commentID, kind).
		Count(&count)

	return int(count), nil
}
