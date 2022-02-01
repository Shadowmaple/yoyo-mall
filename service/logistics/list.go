package logistics

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

type Item struct {
	ID      uint32 `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Status  int8   `json:"status"`
}

// 某订单的物流信息
func GetInfoByOrderID(orderID uint32) (list []*Item, err error) {
	list = make([]*Item, 0)

	records, err := model.GetLogisticsByOrderID(orderID)
	if err != nil {
		return
	}

	for _, record := range records {
		list = append(list, &Item{
			ID:      record.ID,
			Content: record.Content,
			Time:    util.GetStandardTime(record.CreateTime),
			Status:  record.Status,
		})
	}

	return
}

// todo
// 管理端查询物流列表，根据状态筛选
func GetList(limit, page int, state int8) (list []*Item, err error) {
	list = make([]*Item, 0)

	return
}
