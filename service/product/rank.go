package product

type RankItem struct {
	ID          uint32  `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Publisher   string  `json:"publisher"`
	BookName    string  `json:"book_name"`
	Cid         uint32  `json:"cid"`
	Cid2        uint32  `json:"cid2"`
	Price       float32 `json:"price"`
	CurPrice    float32 `json:"cur_price"`
	Image       string  `json:"image"` // 封面图片
	SaleNum     int     `json:"sale_num"`
	CommentNum  int     `json:"comment_num"`
	CommentRate float32 `json:"comment_rate"`
	Score       float32 `json:"score"`
	PublishTime string  `json:"publish_time"`
}

// todo：榜单
func GetRank(kind, limit int, cid, cid2 uint32) (list []*RankItem, err error) {
	list = make([]*RankItem, 0)

	// 暂时先用这个
	filter := FilterItem{Cid: cid, Cid2: cid2, Sort: 0}
	switch kind {
	case 0:
		// 畅销榜
		filter.Sort = 1
	case 1:
		// 新书榜
		filter.Sort = 5
	}

	products, err := List(0, 20, 0, filter)
	if err != nil {
		return
	}

	for _, item := range products {
		list = append(list, &RankItem{
			ID:          item.ID,
			Title:       item.Title,
			Author:      item.Author,
			Publisher:   item.Publisher,
			BookName:    item.BookName,
			Cid:         item.Cid,
			Cid2:        item.Cid2,
			Price:       item.Price,
			CurPrice:    item.CurPrice,
			Image:       item.Image,
			SaleNum:     item.SaleNum,
			CommentNum:  item.CommentNum,
			CommentRate: item.CommentRate,
			Score:       item.Score,
			PublishTime: item.PublishTime,
		})
	}

	return
}
