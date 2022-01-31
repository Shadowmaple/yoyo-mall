package product

import (
	"fmt"
	"log"
	"yoyo-mall/model"
	"yoyo-mall/util"
)

const SearchSQL = `
select t5.*, ifnull(t4.comment_num, 0) as comment_num,
	ifnull(t4.score, 0) as score, ifnull(t4.comment_rate, 0) as comment_rate
from (
	select id, cid, cid2, title, author, publisher, book_name, price, cur_price, images, publish_time,
		ifnull(t1.sale_num, 0) as sale_num
	from product left join (
		select product_id, sum(num) as sale_num
		from order_product
		group by product_id
	) as t1
	on product.id = t1.product_id
	where product.is_deleted = 0 and %s
) as t5 left join (
	select t2.product_id, t2.num as comment_num, t2.score, t3.num/t2.num as comment_rate
	from (
		select product_id, num, total_score/num as score
		from (
			select product_id, count(id) as num, sum(score) as total_score
			from evaluation
			where is_deleted = 0
			group by product_id
		) as t2_1
	) as t2 inner join (
		select product_id, count(id) as num
		from evaluation
		where is_deleted = 0 and level = 0
		group by product_id
	) as t3
	where t2.product_id = t3.product_id
) as t4
on t5.id = t4.product_id
limit %d
offset %d;
`

type SearchItem struct {
	Title     string
	Book      string
	Author    string
	Publisher string
}

func Search(userID uint32, limit, page int, filter SearchItem) (list []*ProductItem, err error) {
	list = make([]*ProductItem, 0)

	filterSQL := "1 = 1"
	if len(filter.Title) > 0 {
		filterSQL += ` and title like '%` + filter.Title + `%'`
	}
	if len(filter.Author) > 0 {
		filterSQL += ` and author like '%` + filter.Author + `%'`
	}
	if len(filter.Book) > 0 {
		filterSQL += ` and book_name like '%` + filter.Book + `%'`
	}
	if len(filter.Publisher) > 0 {
		filterSQL += ` and publisher like '%` + filter.Publisher + `%'`
	}
	log.Println("search sql:", filterSQL)

	offset := limit * page
	querySQL := fmt.Sprintf(SearchSQL, filterSQL, limit, offset)

	models, err := model.ProductSearch(querySQL)
	if err != nil {
		return nil, err
	}

	for _, item := range models {
		hasStar, hasInCart := false, false
		if userID > 0 {
			hasStar = model.HasStar(userID, item.ID)
			hasInCart = model.HasInCart(userID, item.ID)
		}
		image := util.GetFirstImage(item.Images)
		publishTime, _ := util.FormatTime(item.PublishTime)

		list = append(list, &ProductItem{
			ID:          item.ID,
			Title:       item.Title,
			Author:      item.Author,
			Publisher:   item.Publisher,
			BookName:    item.BookName,
			Cid:         item.Cid,
			Cid2:        item.Cid2,
			Price:       item.Price,
			CurPrice:    item.CurPrice,
			Image:       image,
			SaleNum:     item.SaleNum,
			CommentNum:  item.CommentNum,
			CommentRate: item.CommentRate,
			Score:       item.Score,
			PublishTime: publishTime,
			HasStar:     hasStar,
			HasInCart:   hasInCart,
		})
	}

	return
}
