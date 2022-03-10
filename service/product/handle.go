package product

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/log"
	"yoyo-mall/util"
)

type HandleItem struct {
	ID          uint32   `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Publisher   string   `json:"publisher"`
	BookName    string   `json:"book_name"`
	Cid         uint32   `json:"cid"`
	Cid2        uint32   `json:"cid2"`
	Price       float32  `json:"price"`
	CurPrice    float32  `json:"cur_price"`
	Stock       int      `json:"stock"` // 库存
	Images      []string `json:"images"`
	Status      int8     `json:"status"` // 状态，0正常，1下架
	PublishTime string   `json:"publish_time"`
	Detail      string   `json:"detail"`
}

func New(item *HandleItem) (err error) {
	// 权宜之计
	if item.PublishTime == "" {
		item.PublishTime = util.GetStandardTime(util.GetCurrentTime())
	}
	publishTime, err := util.ParseTime(item.PublishTime)
	if err != nil {
		log.Error("parse publishTime error: " + err.Error())
		return
	}

	record := &model.ProductModel{
		Cid:         item.Cid,
		Cid2:        item.Cid2,
		Title:       item.Title,
		Author:      item.Author,
		Publisher:   item.Publisher,
		BookName:    item.BookName,
		Price:       item.Price,
		CurPrice:    item.CurPrice,
		Stock:       item.Stock,
		Status:      item.Status,
		Images:      util.MergeMultiImage(item.Images),
		PublishTime: publishTime,
		Detail:      item.Detail,
	}

	if err = record.Create(); err != nil {
		log.Error("create a new product record error: " + err.Error())
		return
	}

	return
}

func Update(item *HandleItem) (err error) {
	record, err := model.GetProductByID(item.ID)
	if err != nil {
		if err == errno.ErrRecordNotFound {
			err = errno.ErrProductNotExist
			return
		}
		return
	}

	record.Title = item.Title
	record.Author = item.Author
	record.Publisher = item.Publisher
	// record.Cid = item.Cid
	// record.Cid2 = item.Cid2
	// record.BookName = item.BookName
	record.Price = item.Price
	record.CurPrice = item.CurPrice
	// todo:
	// record.Stock = item.Stock
	// record.Images = util.MergeMultiImage(item.Images)
	// record.Status = item.Status
	// record.Detail = item.Detail
	// record.PublishTime, err = util.ParseTime(item.PublishTime)
	// if err != nil {
	// 	log.Error("parse publishTime error: " + err.Error())
	// 	return
	// }

	if err = record.Save(); err != nil {
		log.Error("save product record to DB error: " + err.Error())
		return
	}

	return
}

func Delete(id uint32) (err error) {
	return model.DeleteProduct(id)
}
