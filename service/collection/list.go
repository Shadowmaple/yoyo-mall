package collection

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

type CollectItem struct {
	ID         uint32  `json:"id"`
	Title      string  `json:"title"`
	Author     string  `json:"author"`
	Publisher  string  `json:"publisher"`
	BookName   string  `json:"book_name"`
	Price      float32 `json:"price"`
	CurPrice   float32 `json:"cur_price"`
	Image      string  `json:"image"`
	Stock      int     `json:"stock"`
	CreateTime string  `json:"create_time"`
}

// todo: 一次sql
func List(userID uint32, limit, page int) (list []*CollectItem, err error) {
	list = make([]*CollectItem, 0)
	records, err := model.GetCollection(userID, limit, limit*page)
	if err != nil {
		return
	}
	for _, item := range records {
		product, err1 := model.GetProductByID(item.ProductID)
		if err1 != nil {
			err = err1
			return
		}
		list = append(list, &CollectItem{
			ID:         item.ID,
			Title:      product.Title,
			Author:     product.Author,
			Publisher:  product.Publisher,
			BookName:   product.BookName,
			Price:      product.Price,
			CurPrice:   product.CurPrice,
			Image:      util.GetFirstImage(product.Images),
			Stock:      product.Stock,
			CreateTime: util.GetStandardTime(item.CreateTime),
		})
	}

	return
}
