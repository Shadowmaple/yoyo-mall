package evaluation

type BasicItem struct {
	ID         uint32   `json:"id"`
	OrderID    uint32   `json:"order_id"`
	Content    string   `json:"content"`
	Score      int8     `json:"score"`
	Rank       int8     `json:"rank"`
	IsAnoymous bool     `json:"is_anoymous"`
	Pictures   []string `json:"pictures"`
}

type EvaluationItem struct {
	BasicItem
	Time         string `json:"time"`
	UserNickname string `json:"user_nickname"`
	UserAvatar   string `json:"user_avatar"`
	LikeNum      int    `json:"like_num"`
	ReplyNum     int    `json:"reply_num"`
	HasLiked     bool   `json:"has_liked"`
	CanHandle    bool   `json:"can_handle"`
}
