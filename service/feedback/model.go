package feedback

type BasicItem struct {
	Kind     int8     `json:"kind"`
	Content  string   `json:"content"`
	Pictures []string `json:"pictures"`
}

type FeedbackItem struct {
	BasicItem
	ID           uint32 `json:"id"`
	HasRead      bool   `json:"has_read"`
	UserNickname string `json:"user_nickname"`
	UserAvatar   string `json:"user_avatar"`
}
