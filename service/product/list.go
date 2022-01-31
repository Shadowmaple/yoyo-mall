package product

import (
	"fmt"
	"log"
	"strconv"
	"yoyo-mall/model"
	"yoyo-mall/util"
)

type FilterItem struct {
	Cid  uint32
	Cid2 uint32
	Sort int // 排序类型，0->默认，1->销量，2->价格升序，3->价格降序，4->好评率，5->出版时间降序，6出版时间升序
}

type ProductItem struct {
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
	HasStar     bool    `json:"has_star"`
	HasInCart   bool    `json:"has_in_cart"`
}

/*

商品按类目过滤
select * from product
where is_deleted = 0 and cid = ? and cid2 = ?

涉及的表：
product, order_product, evaluation

销量：
(
select product_id, sum(num) as sale_num
from order_product
group by product_id
) as t1

评论数、平均评分
(
select product_id, num, total_score/num as score
from (
	select product_id, count(id) as num, sum(score) as total_score
	from evaluation
	where is_deleted = 0
	group by product_id
) as t2_1
) as t2

好评数
(
select product_id, count(id) as num
from evaluation
where is_deleted = 0 and level = 0
group by product_id
) as t3

表连接，获取评论数、平均评分、好评率
(
select t2.product_id, t2.num as comment_num, t2.score, t3.num/t2.num as comment_rate
from t2 inner join t3
where t2.product_id = t3.product_id
) as t4

先获取商品信息和销量
(
select id, cid, cid2, title, author, publisher, book_name, price, cur_price, images,
	ifnull(t1.sale_num, 0) as sale_num
from product left join (
	select product_id, sum(num) as sale_num
	from order_product
	group by product_id
) as t1
on product.id = t1.product_id
where product.is_deleted = 0
) as t5


再综合得到商品评价信息：
select t5.*, ifnull(t4.comment_num, 0) as comment_num, ifnull(t4.score, 0) as score, ifnull(t4.comment_rate, 0) as comment_rate
from t5 left join t4
on t5.id = t4.product_id

综合结果：

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
	where product.is_deleted = 0
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

*/

const listQuerySQL = `
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
order by %s
limit %d
offset %d;
`

func List(userID uint32, limit, page int, filter FilterItem) (list []*ProductItem, err error) {
	list = make([]*ProductItem, 0)

	orderByParam := "id"
	switch filter.Sort {
	case 2:
		orderByParam = "cur_price"
	case 3:
		orderByParam = "cur_price desc"
	case 4:
		orderByParam = "comment_rate"
	case 5:
		orderByParam = "publish_time desc"
	case 6:
		orderByParam = "publish_time"
	}

	filterSQL := "1 = 1"
	if filter.Cid > 0 {
		filterSQL += " and cid = " + strconv.Itoa(int(filter.Cid))
	}
	if filter.Cid2 > 0 {
		filterSQL += " and cid2 = " + strconv.Itoa(int(filter.Cid2))
	}
	log.Println("list sql:", orderByParam, filterSQL)

	offset := limit * page
	querySQL := fmt.Sprintf(listQuerySQL, filterSQL, orderByParam, limit, offset)

	models, err := model.ProductList(querySQL)
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
