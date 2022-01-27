package cart

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

type CartItem struct {
	ID       uint32  `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
	CurPrice float32 `json:"cur_price"`
	Image    string  `json:"image"`
	Num      int     `json:"num"`
}

// todo: 一次SQL查询
func List(userID uint32) (list []*CartItem, err error) {
	list = make([]*CartItem, 0)

	records, err := model.GetCarts(userID)
	if err != nil {
		if err == errno.ErrRecordNotFound {
			err = nil
			return
		}
		return
	}

	for _, item := range records {
		product, err1 := model.GetProductByID(item.ProductID)
		if err1 != nil {
			if err == errno.ErrRecordNotFound {
				err = nil
				return
			}
			err = err1
			return
		}

		list = append(list, &CartItem{
			ID:       item.ID,
			Title:    product.Title,
			Author:   product.Author,
			Price:    product.Price,
			CurPrice: product.CurPrice,
			Image:    util.GetFirstImage(product.Images),
			Num:      item.Num,
		})
	}

	return
}
