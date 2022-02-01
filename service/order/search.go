package order

import (
	"fmt"
	"yoyo-mall/model"
)

/*

select * from
order, order_product, product
where order.user_id = ? and order.id = order_product.order_id and order_product.product_id = product.id
and product.title like '%?%'


select t1.*
from product inner join
(
	select order.id, order.status, order.total_fee,
		op.num as product_num, op.product_id, op.image
	from order left join order_product as op
	on order.id = op.order_id
	where order.user_id = ?
) as t1
where product.id = t1.product_id
and product.title like '%?%'


-- 最终决定这么写：
select *
from order
where user_id = ?
and id in (
	select distinct op.order_id
	from order_product as op left join product
	on op.product_id = product.id
	where product.title like '%?%'
)

*/

// 表名要用反引号包起来，不然和MySQL关键字冲突会报错
const SearchSQL = `
select *
from ` + "`order`" + `
where user_id = %d
and id in (
	select distinct op.order_id
	from order_product as op left join product
	on op.product_id = product.id
	where product.title like '%s'
)
`

func Search(userID uint32, key string) (list []*OrderItem, err error) {
	list = make([]*OrderItem, 0)
	searchQuery := fmt.Sprintf(SearchSQL, userID, "%"+key+"%")

	records, err := model.OrderSearch(searchQuery)
	if err != nil {
		return
	}

	var item *OrderItem
	for _, record := range records {
		item, err = processRawOrder(record)
		if err != nil {
			return
		}

		list = append(list, item)
	}

	return
}
