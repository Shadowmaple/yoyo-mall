package logistics

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

// 新物流记录
func NewLogistics(orderID uint32, status int8) error {
	content := ""
	switch status {
	case 0:
		content = "已付款"
	case 1:
		content = "商品已发货"
	case 2:
		content = "商品已签收"
	case 3:
		content = "订单已取消"
	case 4:
		content = "商品正在退货中"
	case 5:
		content = "退货商品成功抵达"
	}

	record := &model.LogisticsModel{
		OrderID:    orderID,
		Status:     status,
		Content:    content,
		CreateTime: util.GetCurrentTime(),
	}

	if err := record.Create(); err != nil {
		return err
	}

	return nil
}
