package product

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

type ProductProfile struct {
	ID          uint32            `json:"id"`
	Title       string            `json:"title"`
	Author      string            `json:"author"`
	Publisher   string            `json:"publisher"`
	BookName    string            `json:"book_name"`
	Cid         uint32            `json:"cid"`
	Cid2        uint32            `json:"cid2"`
	Price       float32           `json:"price"`
	CurPrice    float32           `json:"cur_price"`
	Discount    float32           `json:"discount"`
	Images      []string          `json:"images"`
	CommentList []*EvaluationItem `json:"comment_list"`
	Detail      string            `json:"detail"`
	HasStar     bool              `json:"has_star"`
	HasInCart   bool              `json:"has_in_cart"`
	CartNum     int               `json:"cart_num"`
}

type EvaluationItem struct {
	ID           uint32 `json:"id"`
	UserAvatar   string `json:"user_avatar"`
	UserNickname string `json:"user_nickname"`
	Content      string `json:"content"`
	Score        int8   `json:"score"`
	Time         string `json:"time"`
	ReplyNum     int    `json:"reply_num"`
	HasLike      bool   `json:"has_like"`
}

func GetProfile(id, userID uint32, commentLimit int) (list *ProductProfile, err error) {
	product, err := model.GetProductByID(id)
	if err != nil {
		return
	}

	evaluations, err := model.GetEvaluationList(0, 0, id, commentLimit, 0)
	if err != nil {
		return
	}

	// to do: 评论查询，一次sql，表连接
	evaluationList := make([]*EvaluationItem, 0, 0)
	for _, item := range evaluations {
		user, err1 := model.GetUserByID(item.UserID)
		if err1 != nil {
			err = err1
			return
		}
		replyNum := model.CountComment(item.ID)
		hasLiked := false
		if userID > 0 {
			hasLiked = model.HasLiked(userID, item.ID, 0)
		}
		evaluationList = append(evaluationList, &EvaluationItem{
			ID:           item.ID,
			UserAvatar:   user.Avatar,
			UserNickname: user.Nickname,
			Content:      item.Content,
			Score:        item.Score,
			Time:         util.GetStandardTime(item.CreateTime),
			ReplyNum:     replyNum,
			HasLike:      hasLiked,
		})
	}

	hasStar, hasInCart := false, false
	cartNum := 0
	if userID > 0 {
		hasStar = model.HasStar(userID, id)
		cartNum := model.GetProductNumInCart(userID, id)
		if cartNum > 0 {
			hasInCart = true
		}
	}

	list = &ProductProfile{
		ID:          id,
		Title:       product.Title,
		Author:      product.Author,
		Publisher:   product.Publisher,
		BookName:    product.BookName,
		Cid:         product.Cid,
		Cid2:        product.Cid2,
		Price:       product.Price,
		CurPrice:    product.CurPrice,
		Discount:    product.CurPrice / product.Price,
		Images:      util.ParseMultiImage(product.Images),
		CommentList: evaluationList,
		Detail:      product.Detail,
		HasStar:     hasStar,
		HasInCart:   hasInCart,
		CartNum:     cartNum,
	}

	return
}
