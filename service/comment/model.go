package comment

type BasicItem struct {
	ID         uint32 `json:"id"`
	Content    string `json:"content"`
	IsAnoymous bool   `json:"is_anoymous"`
}

type CommentItem struct {
	BasicItem
	UserNickname string `json:"user_nickname"`
	UserAvatar   string `json:"user_avatar"`
	LikeNum      int    `json:"like_num"`
	HasLiked     bool   `json:"has_liked"`
	CanHandle    bool   `json:"can_handle"`
}
